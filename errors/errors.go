package errors

import (
	base "errors"

	"github.com/delineateio/mimas/log"
)

// DefaultEvent is used if event not provided
const DefaultEvent = "error.event.unknown"

// Errors is a list of errors
type Errors struct {
	Items []error
}

// NewErrors returns the new error wrapper
func NewErrors() *Errors {
	return &Errors{}
}

// Create a new error and log
func (e *Errors) Create(event, desc string) {
	e.Add(event, base.New(desc))
}

// Add errors to the list
func (e *Errors) Add(event string, err error) {
	if err == nil {
		return
	}
	if event == "" {
		event = DefaultEvent
	}
	log.Error(event, err)
	e.Items = append(e.Items, err)
}

// HasErrors indicates if errors exist
func (e *Errors) HasErrors() bool {
	return len(e.Items) > 0
}
