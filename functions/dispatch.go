package functions

import (
	"net/http"

	log "github.com/delineateio/mimas/log"
	messages "github.com/delineateio/mimas/messages"
)

func dispatch(w http.ResponseWriter, r *http.Request, command messages.Command) {
	// Gets the request and binds
	request := messages.NewRequest(r.Method, r.Header)
	binding := messages.NewBinding()
	err := binding.Bind(r, request.Body)
	if err != nil {
		log.Error("request.bind.error", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response := messages.NewJSONResponse()
		command(request, response)

		body := response.ToJSON()
		w.WriteHeader(response.Code)
		for key, value := range response.Headers {
			w.Header().Add(key, value)
		}
		if response.Body != nil {
			_, err = w.Write(body)
			if err != nil {
				log.Error("request.write.error", err)
			}
		}
	}
}
