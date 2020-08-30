package functions

import (
	"github.com/delineateio/mimas/health"
	messages "github.com/delineateio/mimas/messages"
)

func addHealthRoutes(routes []messages.Route) []messages.Route {
	// Default routes
	var items = []messages.Route{
		{
			Method: "GET", Path: "/", Handler: health.Healthz,
		},
		{
			Method: "GET", Path: "/healthz", Handler: health.Healthz,
		},
	}
	return append(items, routes...)
}
