package routes

import (
	"testing"

	"github.com/delineateio/mimas/handlers"
	"github.com/stretchr/testify/assert"
)

const method = "GET"
const path = "/path"

func getBaseRoutes() []Route {
	// Default routes
	return []Route{
		{Method: "GET", Path: "/test", Handler: handlers.NullHandler},
	}
}

func TestNewRoute(t *testing.T) {
	route, err := NewRoute(method, path, handlers.NullHandler)
	if assert.NotNil(t, route) {
		assert.Equal(t, route.Method, method)
		assert.Equal(t, route.Path, path)
	}
	assert.Nil(t, err)
}

func TestRouteMissingMethod(t *testing.T) {
	route, err := NewRoute("", path, handlers.NullHandler)
	assert.Nil(t, route)
	assert.NotNil(t, err)
}

func TestRouteMissingPath(t *testing.T) {
	route, err := NewRoute(method, "", handlers.NullHandler)
	assert.Nil(t, route)
	assert.NotNil(t, err)
}

func TestRouteInvalidPath(t *testing.T) {
	route, err := NewRoute(method, "/%/", handlers.NullHandler)
	assert.Nil(t, route)
	assert.NotNil(t, err)
}

func TestRouteMissingHandler(t *testing.T) {
	route, err := NewRoute(method, path, nil)
	assert.Nil(t, route)
	assert.NotNil(t, err)
}

func TestAddHealthRoutesToNil(t *testing.T) {
	items := AddDefaultRoutes(nil)
	assert.NotNil(t, items)
	assert.Equal(t, len(items), 2)
}

func TestAddHealthRoutesToEmpty(t *testing.T) {
	items := AddDefaultRoutes([]Route{})
	assert.NotNil(t, items)
	assert.Equal(t, len(items), 2)
}

func TestAddHealthWithRoutes(t *testing.T) {
	items := getBaseRoutes()
	items = AddDefaultRoutes(items)
	assert.NotNil(t, items)
	assert.Equal(t, len(items), 3)
}
