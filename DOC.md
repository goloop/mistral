# mistral - reference

The full reference for the `mistral` package: the client, the shared `goloop/ai`
model, chat completions (interface and native), streaming, embeddings and
models.

Ukrainian version: **[DOC.UK.md](DOC.UK.md)**.

## Contents

- [Mental model](#mental-model)
- [Creating a client](#creating-a-client)
- [Generate and Stream](#generate-and-stream)
- [Native chat completions](#native-chat-completions)
- [Tools, images and system prompts](#tools-images-and-system-prompts)
- [Embeddings](#embeddings)
- [Models](#models)
- [Options and errors](#options-and-errors)

## Mental model

`mistral.Client` implements `ai.Client`, the provider-agnostic contract from
`github.com/goloop/ai`. The shared `Generate` and `Stream` cover the common
ground - chat with tools, images and streaming - so code written against the
interface runs on any provider.

Provider-specific power lives in native methods: the full `ChatCompletion`
request, embeddings and model listing. The wire format is chat-completions
compatible.

```go
import (
	"github.com/goloop/ai"
	"github.com/goloop/mistral"
)
```

## Creating a client

```go
c := mistral.New(os.Getenv("MISTRAL_API_KEY"))

c = mistral.New(apiKey, mistral.WithTimeout(30*time.Second))
```

The base URL defaults to `https://api.mistral.ai/v1`. Point `WithBaseURL` at any
compatible endpoint to reuse this client against another gateway.

## Generate and Stream

```go
resp, err := c.Generate(ctx, &ai.Request{
	Model:    mistral.ModelLargeLatest,
	System:   "You are concise.",
	Messages: []ai.Message{ai.UserText("Name three primary colors.")},
})
resp.Text()
resp.ToolCalls()
resp.Usage
```

`Stream` returns `iter.Seq2[ai.Chunk, error]`: text deltas as chunks with
`Text`, a finished tool call as a chunk with `ToolCall`, and a final chunk with
`Done` and `Usage`.

## Native chat completions

For provider-only options build a `ChatRequest` and call `ChatCompletion` or
`ChatCompletionStream`:

```go
resp, err := c.ChatCompletion(ctx, &mistral.ChatRequest{
	Model:          mistral.ModelLargeLatest,
	Messages:       []mistral.ChatMessage{{Role: "user", Content: "as JSON"}},
	ResponseFormat: json.RawMessage(`{"type":"json_object"}`),
})
```

## Tools, images and system prompts

Tool use, images and system prompts use the shared `ai` types: `ai.Tool`,
`ai.Image`, `ai.ToolResult` and a `RoleSystem` message or the `System` field.
Tool results are sent back as `RoleTool` messages whose `ai.ToolResult.ID`
matches the `ai.ToolUse.ID`. Inline image bytes are sent as a base64 data URI
(use a vision model such as `ModelPixtral`).

## Embeddings

```go
vecs, err := c.Embed(ctx, mistral.ModelEmbed, "hello", "world")
// or the full request:
resp, err := c.Embeddings(ctx, &mistral.EmbeddingRequest{
	Model: mistral.ModelEmbed, Input: []string{"hello"},
})
```

## Models

```go
models, err := c.Models(ctx)
m, err := c.GetModel(ctx, mistral.ModelLargeLatest)
```

## Options and errors

Options: `WithBaseURL`, `WithHTTPClient`, `WithTimeout`, `WithMaxRetries`,
`WithHeader`.

A non-success response becomes an `*ai.APIError` with `Status`, `Type`, `Code`,
`Message` and the raw body:

```go
var apiErr *ai.APIError
if errors.As(err, &apiErr) && apiErr.Status == http.StatusTooManyRequests {
	// back off
}
```

Requests missing a model or messages fail before the network with
`ai.ErrNoModel` or `ai.ErrNoMessages`.
