package handlers

import (
	"net/http"

	"github.com/delineateio/mimas/errors"
	"github.com/delineateio/mimas/messages"
)

// NullHandler is a handler that can be used for testing purposes
func NullHandler(request *messages.Request, response *messages.Response) {
	errs := errors.Errors{}
	if request == nil {
		errs.Create("handler.request.error.", "no request provided")
	}
	if errs.HasErrors() {
		response.Code = http.StatusBadRequest
	} else {
		response.Code = http.StatusOK
	}
}
