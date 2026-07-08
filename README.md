[![deps.dev](https://img.shields.io/badge/deps.dev-insights-4c8dbc)](https://deps.dev/go/github.com%2Fgoloop%2Fmistral) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/mistral/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://pkg.go.dev/github.com/goloop/mistral) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)


# mistral

`mistral` is a Go client for the Mistral API. It implements the
`github.com/goloop/ai` interface, so it looks and works like every other goloop
AI provider, and exposes Mistral's native endpoints with their full options on
top.

## Features

- Chat completions: `Generate` for a single response, `Stream` for
  token-by-token output through `iter.Seq2`.
- Tool use (function calling), multimodal image input and system prompts.
- Native `ChatCompletion` and `ChatCompletionStream` with the full option set.
- Embeddings and model listing.
- Retries on 429 and 5xx with backoff; normalized, typed API errors.
- Depends only on `github.com/goloop/ai` and the standard library.

## Installation

```sh
go get github.com/goloop/mistral
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/goloop/ai"
	"github.com/goloop/mistral"
)

func main() {
	c := mistral.New(os.Getenv("MISTRAL_API_KEY"))

	resp, err := c.Generate(context.Background(), &ai.Request{
		Model:    mistral.ModelLargeLatest,
		Messages: []ai.Message{ai.UserText("Say hello in one word.")},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Text())
}
```

## Streaming

```go
for chunk, err := range c.Stream(ctx, req) {
	if err != nil {
		break
	}
	fmt.Print(chunk.Text)
}
```

## Embeddings

```go
vecs, err := c.Embed(ctx, mistral.ModelEmbed, "hello", "world")
```

## Documentation

Full reference: **[DOC.md](DOC.md)** (Ukrainian: **[DOC.UK.md](DOC.UK.md)**).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT - see [LICENSE](LICENSE).
