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
// # Asking what this driver can do
//
// Capabilities describes this driver for the decision taken before a call:
// whether to offer a feature at all, and whether it needs one request or two.
//
//	if ai.SupportsHosted(c, ai.Hosted{Kind: ai.HostedWebSearch}) { ... }
//
// It is a hint and not a permission - support also depends on the model, the
// account and the region - so ai.ErrNoHosted and ai.ErrNoFormat remain the
// source of truth and a caller still handles them. What changes is that a
// refusal the provider only reports as a 400 now arrives as those same
// sentinels, wrapped around the original ai.APIError, so one errors.Is covers
// a limitation this driver knew in advance and one it learned over the wire.
//
// It depends only on goloop/ai and the standard library.
package mistral
