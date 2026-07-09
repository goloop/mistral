package mistral

import (
	"encoding/json"
	"strings"

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
			Message string          `json:"message"`
			Type    string          `json:"type"`
			Code    json.RawMessage `json:"code"`
		} `json:"error"`
	}
	if json.Unmarshal(body, &nested) == nil && nested.Error.Message != "" {
		e.Message = nested.Error.Message
		e.Type = nested.Error.Type
		e.Code = rawToString(nested.Error.Code)
		return e
	}

	var flat struct {
		Message string          `json:"message"`
		Type    string          `json:"type"`
		Code    json.RawMessage `json:"code"`
	}
	if json.Unmarshal(body, &flat) == nil {
		e.Message = flat.Message
		e.Type = flat.Type
		e.Code = rawToString(flat.Code)
	}
	return e
}

// rawToString renders a JSON value that may be a string or a number (some
// gateways send a numeric "code") as a plain string.
func rawToString(r json.RawMessage) string {
	s := strings.TrimSpace(string(r))
	if s == "" || s == "null" {
		return ""
	}
	if s[0] == '"' {
		var str string
		if json.Unmarshal(r, &str) == nil {
			return str
		}
	}
	return s
}
