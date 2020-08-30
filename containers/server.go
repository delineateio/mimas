package containers

import (
	"fmt"
	"os"
	"strings"

	config "github.com/delineateio/mimas/config"
	db "github.com/delineateio/mimas/database"
	log "github.com/delineateio/mimas/log"
	messages "github.com/delineateio/mimas/messages"
	"github.com/fsnotify/fsnotify"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
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
	routes       []messages.Route
}

// NewServer creates a new server
func NewServer(routes []messages.Route) *Server {
	// Gets env
	env := os.Getenv("DIO_ENV")
	location := os.Getenv("DIO_LOCATION")

	server := &Server{
		Env:      env,
		Location: location,
		routes:   routes,
	}
	server.Configure()
	return server
}

// Configure returns the router that will be returned
func (s *Server) Configure() {
	config.NewConfigurator(s.Env, s.Location)
	s.Configurator.LoadWithCallback(s.Reload)
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

// Reload callback for if the server is restarted
func (s *Server) Reload(in fsnotify.Event) {
	s.setLogger()
	s.setMode()
	s.setTimeOuts()

	config.NewConfigurator(s.Env, s.Location).LoadWithCallback(s.Reload)
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
