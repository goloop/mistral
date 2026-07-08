package mistral

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/goloop/ai"
)

func TestChatCompletionNative(t *testing.T) {
	c, done := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"model":"m","choices":[{"index":0,`+
			`"message":{"role":"assistant","content":"hi"},"finish_reason":"stop"}]}`)
	})
	defer done()

	resp, err := c.ChatCompletion(context.Background(), &ChatRequest{
		Model:    "m",
		Messages: []ChatMessage{{Role: "user", Content: "hi"}},
	})
	if err != nil || resp.Choices[0].Message.Content != "hi" {
		t.Fatalf("chat: %v %+v", err, resp)
	}
}

func TestEmbed(t *testing.T) {
	c, done := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/embeddings") {
			t.Errorf("path = %q", r.URL.Path)
		}
		io.WriteString(w, `{"model":"mistral-embed","data":[`+
			`{"index":1,"embedding":[0.3,0.4]},{"index":0,"embedding":[0.1,0.2]}]}`)
	})
	defer done()

	vecs, err := c.Embed(context.Background(), ModelEmbed, "a", "b")
	if err != nil {
		t.Fatal(err)
	}
	if len(vecs) != 2 || vecs[0][0] != 0.1 || vecs[1][1] != 0.4 {
		t.Errorf("vecs = %+v", vecs)
	}
}

func TestModels(t *testing.T) {
	c, done := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/models/mistral-large-latest"):
			io.WriteString(w, `{"id":"mistral-large-latest","object":"model"}`)
		default:
			io.WriteString(w, `{"data":[{"id":"mistral-large-latest","object":"model"}]}`)
		}
	})
	defer done()

	ctx := context.Background()
	if models, err := c.Models(ctx); err != nil || len(models) != 1 {
		t.Fatalf("models: %v %+v", err, models)
	}
	m, err := c.GetModel(ctx, "mistral-large-latest")
	if err != nil || m.ID != "mistral-large-latest" {
		t.Fatalf("get model: %v %+v", err, m)
	}
}

func TestErrorFlat(t *testing.T) {
	c, done := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message":"Unauthorized","type":"invalid_request_error"}`)
	})
	defer done()

	_, err := c.Generate(context.Background(), &ai.Request{
		Model: "m", Messages: []ai.Message{ai.UserText("hi")},
	})
	var apiErr *ai.APIError
	if !errors.As(err, &apiErr) || apiErr.Status != 401 || apiErr.Message != "Unauthorized" {
		t.Fatalf("err = %v", err)
	}
}
