package mistral_test

import (
	"encoding/json"
	"fmt"

	"github.com/goloop/ai"
	"github.com/goloop/mistral"
)

func ExampleNew() {
	c := mistral.New("...")
	_ = c // use c.Generate, c.Stream, c.ChatCompletion, ...
	fmt.Println(mistral.ModelLargeLatest)
	// Output: mistral-large-latest
}

// ExampleClient_Generate builds a request. Sending it needs a real API key, so
// this example only shows the shape.
func ExampleClient_Generate() {
	req := &ai.Request{
		Model: mistral.ModelLargeLatest,
		Messages: []ai.Message{
			ai.UserText("Name the capital of France."),
		},
	}
	fmt.Println(req.Model, len(req.Messages))
	// Output: mistral-large-latest 1
}

// ExampleClient_FIM shows a fill-in-the-middle request. The model completes the
// gap between Prompt and Suffix; read the inserted text from the first choice.
func ExampleClient_FIM() {
	req := &mistral.FIMRequest{
		Model:  mistral.ModelCodestral,
		Prompt: "def add(a, b):\n    ",
		Suffix: "\n    return result",
	}
	fmt.Println(req.Model)
	// Output: codestral-latest
}

// ExampleTool shows a tool definition passed with a request.
func ExampleTool() {
	tool := ai.Tool{
		Name:        "get_weather",
		Description: "Get the current weather for a city.",
		Schema:      json.RawMessage(`{"type":"object","properties":{"city":{"type":"string"}}}`),
	}
	fmt.Println(tool.Name)
	// Output: get_weather
}
