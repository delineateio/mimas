package handlers

import (
	"net/http"

	"github.com/delineateio/mimas/msgs"
)

// HealthzHandler is the health check - the name is inspired by
// a forgotten source that this is the naming conventions at Google
func HealthzHandler(request *msgs.Request, response *msgs.Response) {
	response.Code = http.StatusOK
}
