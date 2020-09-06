package handlers

import (
	"net/http"

	"github.com/delineateio/mimas/messages"
)

// HealthzHandler is the health check - the name is inspired by
// a forgotten source that this is the naming conventions at Google
func HealthzHandler(request *messages.Request, response *messages.Response) {
	response.Code = http.StatusOK
}
