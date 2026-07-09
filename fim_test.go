package mistral

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestFIM(t *testing.T) {
	c, done := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/fim/completions" {
			t.Errorf("path = %s", r.URL.Path)
		}
		var req FIMRequest
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &req); err != nil {
			t.Fatal(err)
		}
		if req.Prompt != "def add(a, b):" || req.Suffix != "return result" {
			t.Errorf("req = %+v", req)
		}
		io.WriteString(w, `{"id":"1","model":"codestral-latest","choices":[{"index":0,`+
			`"message":{"role":"assistant","content":"\n    result = a + b\n    "},`+
			`"finish_reason":"stop"}],"usage":{"prompt_tokens":8,"completion_tokens":6}}`)
	})
	defer done()

	resp, err := c.FIM(context.Background(), &FIMRequest{
		Model:  ModelCodestral,
		Prompt: "def add(a, b):",
		Suffix: "return result",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Choices) != 1 || contentText(resp.Choices[0].Message.Content) == "" {
		t.Errorf("resp = %+v", resp)
	}
	if resp.Usage.CompletionTokens != 6 {
		t.Errorf("usage = %+v", resp.Usage)
	}
}

func TestFIMError(t *testing.T) {
	c, done := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, `{"message":"invalid prompt"}`)
	})
	defer done()

	_, err := c.FIM(context.Background(), &FIMRequest{Model: "m", Prompt: "x"})
	if err == nil {
		t.Fatal("want error")
	}
}
