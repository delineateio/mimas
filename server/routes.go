package server

import (
	health "github.com/delineateio/mimas/health"
	messages "github.com/delineateio/mimas/messages"
	"github.com/gin-gonic/gin"
)

func convertRoutes(routes []messages.Route) []gin.RouteInfo {
	// Default routes
	var items = []gin.RouteInfo{
		{
			Method: "GET", Path: "/", HandlerFunc: func(ctx *gin.Context) {
				Dispatch(ctx, health.Healthz)
			},
		},
		{
			Method: "GET", Path: "/healthz", HandlerFunc: func(ctx *gin.Context) {
				Dispatch(ctx, health.Healthz)
			},
		},
	}
	// Adds the user defined routes
	for _, route := range routes {
		items = append(items, getGinRoute(route))
	}

	return items
}

func getGinRoute(route messages.Route) gin.RouteInfo {
	return gin.RouteInfo{
		Method: route.Method, Path: route.Path, HandlerFunc: func(ctx *gin.Context) {
			Dispatch(ctx, route.Handler)
		},
	}
}
