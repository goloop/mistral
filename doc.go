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
// # Hosted capabilities
//
// This provider's hosted search belongs to its conversations and agents
// endpoints, not to the chat endpoint this package speaks: a chat reply
// carries no references, and references are what make a searched answer
// checkable. ai.Hosted is refused with ai.ErrNoHosted.
//
// The refusal is the documented behavior, not a gap waiting to be filled
// in silence: an answer produced without the search that was asked for
// looks exactly like one produced with it. A caller who would rather have
// the answer anyway asks again without ai.Request.Hosted.
//
// It depends only on goloop/ai and the standard library.
package mistral
