package common

import (
	"github.com/mitchellh/mapstructure"
)

// Command performs the required action for the service
type Command func(request *Request, response *Response)

// Request generically represents inputs to the service
type Request struct {
	Body map[string]interface{}
}

// Response generically represents outputs from the service
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body,omitempty"`
}

// Map maps request body to the domain model
func (r *Request) Map(entity interface{}) error {
	err := mapstructure.Decode(r.Body, entity)
	if err != nil {
		Error("request was not understood", err)
	}
	return err
}
