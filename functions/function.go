package functions

import (
	"net/http"

	log "github.com/delineateio/mimas/log"
	messages "github.com/delineateio/mimas/messages"
	"github.com/gorilla/mux"
)

// NewFunction creates a new function
func NewFunction(routes []messages.Route) *mux.Router {
	// Adds the default health routes
	routes = addHealthRoutes(routes)

	router := mux.NewRouter()
	if routes != nil {
		for _, route := range routes {
			current := route
			router.HandleFunc(current.Path, func(w http.ResponseWriter, r *http.Request) {
				dispatch(w, r, current.Handler)
			}).Methods(current.Method)
		}
		log.Info("server.routes.add", "routes have been added")
	}

	return router
}

// -----------------------------------------------------------------------------
var function = NewFunction(getRoutes())

func getRoutes() []messages.Route {
	return nil
}

// F represents cloud function entry point
func F(w http.ResponseWriter, r *http.Request) {
	function.ServeHTTP(w, r)
}

// -----------------------------------------------------------------------------
