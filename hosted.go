package mistral

import (
	"fmt"

	"github.com/goloop/ai"
)

// checkHosted refuses a request that asks the provider to run a capability on
// its own side.
//
// This provider's hosted search belongs to its conversations and agents
// endpoints, not to the chat endpoint this driver speaks. Its own documentation
// gives the reason: a chat reply carries no references, and references are the
// half that makes a searched answer checkable. Supporting it here means
// speaking a second endpoint, which a later version can do.
//
// Returning [ai.ErrNoHosted] before the request leaves is the documented
// behavior, not a placeholder: an answer produced without the search that was
// asked for looks exactly like one produced with it, so failing loudly is the
// only way a caller can tell the difference.
func checkHosted(req *ai.Request) error {
	if len(req.Hosted) == 0 {
		return nil
	}
	return fmt.Errorf("%w: this provider's search lives on another endpoint", ai.ErrNoHosted)
}
