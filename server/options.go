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
		Entities: make([]interface{}, 0),
	}, nil
}

// AddRoute is used to add another route to the
func (o *Options) AddRoute(method, path string, handler handlers.Handler) {
	o.Routes = append(o.Routes, routes.Route{
		Method: method, Path: path, Handler: handler,
	})
}

// AddEntities is used to add the required entities for migration
func (o *Options) AddEntities(entities ...interface{}) {
	o.Entities = append(o.Entities, entities...)
}
