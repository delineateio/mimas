package functions

import (
	"errors"
	"net/http"

	e "github.com/delineateio/mimas/errors"
	"github.com/delineateio/mimas/messages"
)

func dispatch(writer http.ResponseWriter, r *http.Request, command messages.Command) {
	// Gets the request and binds
	errs := e.NewErrors()
	request, err := messages.NewRequest(r.Method, r.Header)
	errs.Add("request.bind.error", err)

	binding := messages.NewBinding()
	err = binding.Bind(r, request.Body)
	errs.Add("request.bind.error", err)

	if errs.HasErrors() {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		response := messages.NewJSONResponse()
		command(request, response)
		writeResponse(writer, response, errs)
	}
}

func writeResponse(w http.ResponseWriter, response *messages.Response, errs *e.Errors) {
	for key, value := range response.Headers {
		w.Header().Add(key, value)
	}
	if response.Body != nil {
		if !response.IsValid() {
			errs.Add("response.body.error", errors.New("invalid response body"))
		}
		if !errs.HasErrors() {
			_, err := w.Write(response.ToBytes())
			errs.Add("response.body.error", err)
		}
	}
	if errs.HasErrors() {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(response.Code)
	}
}
