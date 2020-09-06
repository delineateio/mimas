package msgs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// IBinding for all repositories
type IBinding interface {
	Bind(req *http.Request, obj interface{}) error
}

// JSONBinding that reprents the access to the underlying database
type JSONBinding struct{}

// NewBinding creates the binder of the correct type
func NewBinding() *JSONBinding {
	return &JSONBinding{}
}

// Bind reads the data from request
func (binding *JSONBinding) Bind(req *http.Request, obj interface{}) error {
	if req == nil {
		return fmt.Errorf("invalid request")
	}
	if req.Body == nil {
		return nil
	}
	return binding.decodeJSON(req.Body, obj)
}

func (binding *JSONBinding) decodeJSON(reader io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(obj)
}
