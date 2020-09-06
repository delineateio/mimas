package messages

import (
	"encoding/json"
)

// Response generically represents outputs from the service
type Response struct {
	Headers map[string]string
	Code    int         `json:"code"`
	Body    interface{} `json:"body,omitempty"`
}

// NewJSONResponse creates a new response
func NewJSONResponse() *Response {
	return &Response{
		Headers: addJSONHeaders(),
	}
}

func addJSONHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	return headers
}

// HasBody indicates if the response has a body
func (response *Response) HasBody() bool {
	return response.Body != nil
}

// IsValid indicates if the body is validate
func (response *Response) IsValid() bool {
	input := response.ToBytes()
	if input == nil {
		return true
	}
	var container map[string]interface{}
	err := json.Unmarshal(input, &container)
	return err == nil
}

// ToBytes object to byte array
func (response *Response) ToBytes() []byte {
	if response.Body == nil {
		return nil
	}
	data, err := json.Marshal(response.Body)
	if err != nil {
		return nil
	}
	return data
}
