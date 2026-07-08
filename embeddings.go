package mistral

import "context"

// EmbeddingRequest is the native embeddings request.
type EmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

// Embedding is one embedding vector with its position in the input.
type Embedding struct {
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

// EmbeddingResponse is the native embeddings response.
type EmbeddingResponse struct {
	Model string      `json:"model"`
	Data  []Embedding `json:"data"`
	Usage ChatUsage   `json:"usage"`
}

// Embeddings sends a native embeddings request.
func (c *Client) Embeddings(
	ctx context.Context,
	req *EmbeddingRequest,
) (*EmbeddingResponse, error) {
	var out EmbeddingResponse
	if err := c.postJSON(ctx, "/embeddings", req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Embed embeds one or more inputs and returns their vectors in order.
func (c *Client) Embed(
	ctx context.Context,
	model string,
	input ...string,
) ([][]float64, error) {
	resp, err := c.Embeddings(ctx, &EmbeddingRequest{Model: model, Input: input})
	if err != nil {
		return nil, err
	}
	vecs := make([][]float64, len(resp.Data))
	for _, e := range resp.Data {
		if e.Index >= 0 && e.Index < len(vecs) {
			vecs[e.Index] = e.Embedding
		}
	}
	return vecs, nil
}
