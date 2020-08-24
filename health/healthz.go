package health

import (
	"net/http"

	messages "github.com/delineateio/mimas/messages"
)

// Healthz is the health check - the name is inspired by
// a forgotten source that this is the naming conventions at Google
func Healthz(request *messages.Request, response *messages.Response) {
	status := NewMonitor().GetStatus()

	// If there are no checks configured then the service is good
	// Otherwise the status will be taken from the status check
	if !status.IsMonitoring {
		response.Code = http.StatusOK
		return
	}

	if status.Failed {
		response.Code = http.StatusServiceUnavailable
	} else {
		response.Code = http.StatusOK
	}
}
