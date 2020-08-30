package messages

import (
	"encoding/json"
	"net/http"

	log "github.com/delineateio/mimas/log"
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

// ToJSON object to byte array
func (response *Response) ToJSON() []byte {
	body, err := json.Marshal(response.Body)
	if err != nil {
		log.Error("request.error", err)
		response.Code = http.StatusInternalServerError
	}
	return body
}
