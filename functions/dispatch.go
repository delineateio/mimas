package functions

import (
	"errors"
	"net/http"

	e "github.com/delineateio/mimas/errors"
	"github.com/delineateio/mimas/handlers"
	"github.com/delineateio/mimas/msgs"
)

func dispatch(writer http.ResponseWriter, r *http.Request, handler handlers.Handler) {
	// Gets the request and binds
	errs := e.NewErrors()
	request, err := msgs.NewRequest(r.Method, r.Header)
	errs.Add("request.bind.error", err)

	binding := msgs.NewBinding()
	err = binding.Bind(r, request.Body)
	errs.Add("request.bind.error", err)

	if errs.HasErrors() {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		response := msgs.NewJSONResponse()
		handler(request, response)
		writeResponse(writer, response, errs)
	}
}

func writeResponse(w http.ResponseWriter, response *msgs.Response, errs *e.Errors) {
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
