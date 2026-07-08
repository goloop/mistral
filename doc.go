// Package mistral is a client for the Mistral API, built on the goloop/ai
// interface.
//
// The Client implements ai.Client, so Generate and Stream work the same as
// with any other goloop AI provider. On top of that it exposes the native
// chat completions endpoint with its full options, embeddings and model
// listing. The wire format is chat-completions compatible.
//
//	c := mistral.New(os.Getenv("MISTRAL_API_KEY"))
//	resp, err := c.Generate(ctx, &ai.Request{
//	    Model:    mistral.ModelLargeLatest,
//	    Messages: []ai.Message{ai.UserText("Say hello in one word.")},
//	})
//
// It depends only on goloop/ai and the standard library.
package mistral
