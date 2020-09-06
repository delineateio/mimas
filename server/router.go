package server

import (
	c "github.com/delineateio/mimas/cors"
	"github.com/delineateio/mimas/log"
	"github.com/delineateio/mimas/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CreateRouter returns the router that will be returned
func (s *Server) createRouter() *gin.Engine {
	// Misconfiguration can lead to the service not starting
	// The wrapper func defaults to 'release' if that is the case
	gin.SetMode(s.mode)
	router := gin.Default()
	addCORS(router)

	// Adds healthz at the route
	log.Info("server.router.create", "created the GIN router")
	s.options.Routes = routes.AddDefaultRoutes(s.options.Routes)

	// Adds the routes
	if s.options.Routes != nil {
		for _, route := range convertRoutes(s.options.Routes) {
			router.Handle(route.Method, route.Path, route.HandlerFunc)
		}
		log.Info("server.routes.add", "routes have been added")
	}

	return router
}

func addCORS(router *gin.Engine) {
	// Sets the CORS parameters
	settings := cors.DefaultConfig()
	corsConfig := c.NewCORSConfig()
	settings.AllowOrigins = corsConfig.AllowOrigins
	settings.AllowMethods = corsConfig.AllowMethods
	settings.AllowHeaders = corsConfig.AllowHeaders
	settings.ExposeHeaders = corsConfig.ExposeHeaders
	settings.AllowCredentials = corsConfig.AllowCredentials
	settings.MaxAge = corsConfig.MaxAge
	router.Use(cors.New(settings))
}
