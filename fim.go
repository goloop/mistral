package mistral

import "context"

// FIMRequest is a fill-in-the-middle completion request. The model (a Codestral
// model) completes the gap between Prompt and Suffix; leave Suffix empty for a
// plain prefix completion.
type FIMRequest struct {
	Model       string   `json:"model"`
	Prompt      string   `json:"prompt"`
	Suffix      string   `json:"suffix,omitempty"`
	MaxTokens   int      `json:"max_tokens,omitempty"`
	Temperature *float64 `json:"temperature,omitempty"`
	TopP        *float64 `json:"top_p,omitempty"`
	Stop        []string `json:"stop,omitempty"`
}

// FIM sends a fill-in-the-middle completion request. The response reuses the
// chat completions shape; read the inserted text from the first choice's
// message content.
func (c *Client) FIM(ctx context.Context, req *FIMRequest) (*ChatResponse, error) {
	var out ChatResponse
	if err := c.postJSON(ctx, "/fim/completions", req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
