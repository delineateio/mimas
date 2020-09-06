package routes

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/delineateio/mimas/handlers"
	"github.com/delineateio/mimas/log"
	"github.com/delineateio/mimas/msgs"
)

// Route for handling
type Route struct {
	Method  string
	Path    string
	Handler handlers.Handler
}

// NewRoute creates a new route with validation
func NewRoute(method, path string, handler handlers.Handler) (*Route, error) {
	// Validates the parameters before returning the routes
	err := msgs.ValidateMethod(method)
	if err != nil {
		return nil, err
	}
	result, err := validatePath(path)
	if !result {
		return nil, err
	}
	if handler == nil {
		return nil, errors.New("no route handler was provided")
	}
	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}, nil
}

func validatePath(path string) (bool, error) {
	if path == "" {
		return false, errors.New("no route path was provided")
	}
	_, err := url.Parse(path)
	if err != nil {
		return false, fmt.Errorf("path '%s' was not valid", path)
	}
	return true, nil
}

// AddDefaultRoutes adds the default routes
func AddDefaultRoutes(current []Route) []Route {
	// Default routes
	var items = []Route{
		{Method: "GET", Path: "/", Handler: handlers.HealthzHandler},
		{Method: "GET", Path: "/healthz", Handler: handlers.HealthzHandler},
	}
	log.Info("routes.added", "added the default routes")

	return append(items, current...)
}
