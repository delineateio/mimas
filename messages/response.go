package messages

import (
	log "github.com/delineateio/mimas/log"
	"github.com/mitchellh/mapstructure"
)

// Response generically represents outputs from the service
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body,omitempty"`
}

// Map maps request body to the domain model
func (r *Request) Map(entity interface{}) error {
	err := mapstructure.Decode(r.Body, entity)
	if err != nil {
		log.Error("request was not understood", err)
	}
	return err
}
