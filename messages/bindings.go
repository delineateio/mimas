package messages

import (
	"net/http"
)

// IBinding for all repositories
type IBinding interface {
	Bind(req *http.Request, obj interface{}) error
}

// NewBinding creates the binder of the correct type
func NewBinding() *JSONBinding {
	return &JSONBinding{}
}
