package messages

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JSONBinding that reprents the access to the underlying database
type JSONBinding struct{}

// Bind reads the data from request
func (binding *JSONBinding) Bind(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return binding.decodeJSON(req.Body, obj)
}

func (binding *JSONBinding) decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(obj)
}
