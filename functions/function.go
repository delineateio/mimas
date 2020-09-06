package functions

import (
	"fmt"
	"net/http"

	"github.com/delineateio/mimas/log"
	"github.com/delineateio/mimas/routes"
	"github.com/gorilla/mux"
)

// NewFunction creates a new function
func NewFunction(current []routes.Route) *mux.Router {
	// Adds the default health routes
	current = routes.AddDefaultRoutes(current)
	router := mux.NewRouter()
	if current != nil {
		for _, route := range current {
			current := route
			router.HandleFunc(current.Path, func(w http.ResponseWriter, r *http.Request) { dispatch(w, r, current.Handler) }).Methods(current.Method)
			log.Debug("server.routes.add", fmt.Sprintf("method: %s, path: %s", current.Method, current.Path))
		}
		log.Info("server.routes.add", "all routes have been added")
	}

	return router
}

// -----------------------------------------------------------------------------
/* var function = NewFunction(getRoutes())
func getRoutes() []messages.Route {
	return nil
}
// F represents cloud function entry point
func F(w http.ResponseWriter, r *http.Request) {
	function.ServeHTTP(w, r)
} */
// -----------------------------------------------------------------------------
