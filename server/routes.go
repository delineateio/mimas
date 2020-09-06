package server

import (
	"github.com/delineateio/mimas/routes"
	"github.com/gin-gonic/gin"
)

func convertRoutes(current []routes.Route) []gin.RouteInfo {
	var items = []gin.RouteInfo{}
	for _, route := range current {
		items = append(items, getGinRoute(route))
	}
	return items
}

func getGinRoute(route routes.Route) gin.RouteInfo {
	return gin.RouteInfo{
		Method: route.Method, Path: route.Path, HandlerFunc: func(ctx *gin.Context) {
			dispatch(ctx, route.Handler)
		},
	}
}
