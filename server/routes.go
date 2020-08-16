package server

import (
	c "github.com/delineateio/mimas/common"
	"github.com/gin-gonic/gin"
)

func convertRoutes(routes []c.Route) []gin.RouteInfo {
	// Default routes
	var items = []gin.RouteInfo{
		{
			Method: "GET", Path: "/", HandlerFunc: func(ctx *gin.Context) {
				Dispatch(ctx, c.Healthz)
			},
		},
		{
			Method: "GET", Path: "/healthz", HandlerFunc: func(ctx *gin.Context) {
				Dispatch(ctx, c.Healthz)
			},
		},
	}
	// Adds the user defined routes
	for _, route := range routes {
		items = append(items, getGinRoute(route))
	}

	return items
}

func getGinRoute(route c.Route) gin.RouteInfo {
	return gin.RouteInfo{
		Method: route.Method, Path: route.Path, HandlerFunc: func(ctx *gin.Context) {
			Dispatch(ctx, route.Handler)
		},
	}
}
