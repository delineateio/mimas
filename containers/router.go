package containers

import (
	"time"

	config "github.com/delineateio/mimas/config"
	log "github.com/delineateio/mimas/log"
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

	// Adds the routes
	if s.routes != nil {
		for _, route := range convertRoutes(s.routes) {
			router.Handle(route.Method, route.Path, route.HandlerFunc)
		}
		log.Info("server.routes.add", "routes have been added")
	}

	return router
}

func addCORS(router *gin.Engine) {
	// Sets the CORS parameters
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.GetStrings("server.cors.allow_origins")
	corsConfig.AllowMethods = config.GetStrings("server.cors.allow_methods")
	corsConfig.AllowHeaders = config.GetStrings("server.cors.allow_headers")
	corsConfig.ExposeHeaders = config.GetStrings("server.cors.expose_headers")
	corsConfig.AllowCredentials = config.GetBool("server.cors.allow_credentials", false)
	corsConfig.MaxAge = config.GetDuration("server.cors.allow_credentials", time.Hour)
	router.Use(cors.New(corsConfig))
}
