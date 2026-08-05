// Package mistral is a client for the Mistral API, built on the goloop/ai
// interface.
//
// The Client implements ai.Client, so Generate and Stream work the same as
// with any other goloop AI provider. On top of that it exposes the native
// chat completions endpoint with its full options, embeddings, fill-in-the-
// middle completions (FIM, for Codestral models) and model listing. The wire
// format is chat-completions compatible.
//
//	c := mistral.New(os.Getenv("MISTRAL_API_KEY"))
//	resp, err := c.Generate(ctx, &ai.Request{
//	    Model:    mistral.ModelLargeLatest,
//	    Messages: []ai.Message{ai.UserText("Say hello in one word.")},
//	})
//
// # Structured output
//
// ai.Request.Format maps onto the provider's response_format, so a request for
// JSON is enforced rather than merely asked for, and ai.Response.JSON decodes
// the reply. Plain JSON mode also puts ai.Format.Instruction into the system
// prompt, because this wire format rejects json_object unless the word "json"
// appears in the messages.
//
// It depends only on goloop/ai and the standard library.
package mistral
