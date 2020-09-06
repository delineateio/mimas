package messages

import (
	log "github.com/delineateio/mimas/log"
	"github.com/mitchellh/mapstructure"
)

// NewRequest Creates a new generic request
func NewRequest(method string, headers map[string][]string) (*Request, error) {
	// Validate the method
	err := ValidateMethod(method)
	if err != nil {
		return nil, err
	}
	return &Request{
		Method:  method,
		Headers: headers,
		Body:    make(map[string]interface{}),
	}, nil
}

// Request generically represents inputs to the service
type Request struct {
	Method  string
	Headers map[string][]string
	Body    map[string]interface{}
}

// Translate maps request body to the domain model
func (request *Request) Translate(entity interface{}) error {
	err := mapstructure.Decode(request.Body, entity)
	if err != nil {
		log.Error("request was not understood", err)
	}
	return err
}
