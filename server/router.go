package server

import (
	"time"

	c "github.com/delineateio/mimas/config"
	log "github.com/delineateio/mimas/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CreateRouter returns the router that will be returned
func (s *Server) CreateRouter() *gin.Engine {
	// Misconfiguration can lead to the service not starting
	// The wrapper func defaults to 'release' if that is the case
	gin.SetMode(s.Mode)
	router := gin.Default()
	addCors(router)

	// Adds healthz at the route
	log.Info("server.router.create", "created the GIN router")

	// Adds the routes
	if s.Routes != nil {
		for _, route := range convertRoutes(s.Routes) {
			router.Handle(route.Method, route.Path, route.HandlerFunc)
		}
		log.Info("server.routes.add", "routes have been added")
	}

	return router
}

func addCors(router *gin.Engine) {
	// Sets the CORS parameters
	config := cors.DefaultConfig()
	config.AllowOrigins = c.GetStrings("server.cors.allow_origins")
	config.AllowMethods = c.GetStrings("server.cors.allow_methods")
	config.AllowHeaders = c.GetStrings("server.cors.allow_headers")
	config.ExposeHeaders = c.GetStrings("server.cors.expose_headers")
	config.AllowCredentials = c.GetBool("server.cors.allow_credentials", false)
	config.MaxAge = c.GetDuration("server.cors.allow_credentials", time.Hour)
	router.Use(cors.New(config))
}
