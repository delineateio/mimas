package server

import (
	"github.com/delineateio/mimas/env"
	"github.com/delineateio/mimas/handlers"
	"github.com/delineateio/mimas/routes"
)

// Options is the values that can be tailored
type Options struct {
	Env      string
	Location string
	Routes   []routes.Route
	Entities []interface{}
}

// NewDefaultOptions creates a new default option
func NewDefaultOptions() (*Options, error) {
	vars := env.NewEnv()
	return &Options{
		Env:      vars.Read("DIO_ENV", "io"),
		Location: vars.Read("DIO_LOCATION", "/config"),
		Routes:   []routes.Route{},
	}, nil
}

// Add is used to add another route to the
func (o *Options) Add(method, path string, handler handlers.Handler) {
	o.Routes = append(o.Routes, routes.Route{
		Method: method, Path: path, Handler: handler,
	})
}
