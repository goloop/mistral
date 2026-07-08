package mistral

import (
	"encoding/json"

	"github.com/goloop/ai"
)

// parseError turns a non-success response body into an *ai.APIError. Mistral
// reports errors either nested as {"error":{"message","type","code"}} or flat
// as {"message":"...","type":"...","code":"..."}; both are handled.
func parseError(status int, body []byte) error {
	e := &ai.APIError{
		Status: status,
		Raw:    append(json.RawMessage(nil), body...),
	}

	var nested struct {
		Error struct {
			Message string `json:"message"`
			Type    string `json:"type"`
			Code    string `json:"code"`
		} `json:"error"`
	}
	if json.Unmarshal(body, &nested) == nil && nested.Error.Message != "" {
		e.Message = nested.Error.Message
		e.Type = nested.Error.Type
		e.Code = nested.Error.Code
		return e
	}

	var flat struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	}
	if json.Unmarshal(body, &flat) == nil {
		e.Message = flat.Message
		e.Type = flat.Type
		e.Code = flat.Code
	}
	return e
}
