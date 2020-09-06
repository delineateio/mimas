package server

import (
	"fmt"
	"strings"

	"github.com/delineateio/mimas/config"
	"github.com/delineateio/mimas/db"
	"github.com/delineateio/mimas/environment"
	"github.com/delineateio/mimas/log"
	"github.com/delineateio/mimas/routes"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

// Server represents the encapulsation of a service
// Don't rely on server defaults as this could significant impact performance]
// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
type Server struct {
	Env          string
	Location     string
	Configurator config.Configurator
	Repository   db.IRepository
	mode         string
	router       *gin.Engine
	routes       []routes.Route
}

// NewServer creates a new server
func NewServer(current []routes.Route) *Server {
	// Gets env
	env := environment.NewEnv()
	server := &Server{
		Env:      env.Read("DIO_ENV", "io"),
		Location: env.Read("DIO_LOCATION", "/config"),
		routes:   current,
	}
	server.Configure()
	return server
}

// Configure returns the router that will be returned
func (s *Server) Configure() {
	config.NewConfigurator(s.Env, s.Location, afero.NewOsFs())
	s.setLogger()
	s.setMode()
	s.setTimeOuts()
}

func (s *Server) setLogger() {
	level := config.GetString("logging.level", "warn")
	log.NewLogger(level).Load()
	log.Info("config.initialised", fmt.Sprintf("the env config has been set to '%s'", s.Env))
}

func (s *Server) setMode() {
	mode := strings.ToLower(config.GetString("server.mode", "release"))
	if mode != gin.ReleaseMode && mode != gin.DebugMode {
		log.Warn("server.mode", "Configuration incorrect, defaulted to 'release'")
		mode = gin.ReleaseMode
	}
	s.mode = mode
}

func (s *Server) setTimeOuts() {
	// Sets the timeouts
	timeOuts := newTimeOuts()
	endless.DefaultReadTimeOut = timeOuts.read
	endless.DefaultWriteTimeOut = timeOuts.write
	endless.DefaultHammerTime = timeOuts.hammer
	log.Info("server.timeouts", "server timeout configuration completed")
}

// Listen the server and ensure it's configured
func (s *Server) Listen() {
	// Migrates the database
	err := s.Repository.Migrate()
	if err != nil {
		log.Warn("server.start", "there could be issues as the server did not start cleanly")
	}

	// Creates the server / router
	s.router = s.createRouter()
	port := config.GetString("server.port", "1102")
	_ = endless.ListenAndServe(":"+port, s.router)

	// Confirms that the server has been shutdown
	if err != nil {
		log.Info("server.shutdown", "server has shutdown.  Goodbye!")
	}
}
