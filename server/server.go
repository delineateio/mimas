package server

import (
	"fmt"
	"strings"

	"github.com/delineateio/mimas/config"
	"github.com/delineateio/mimas/db"
	"github.com/delineateio/mimas/log"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

// Server represents the encapulsation of a service
// Don't rely on server defaults as this could significant impact performance]
// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
type Server struct {
	options *Options
	mode    string
	router  *gin.Engine
}

// NewServer creates a new server
func NewServer(opts *Options) *Server {
	server := &Server{
		options: opts,
	}
	config.NewConfigurator(opts.Env, opts.Location, afero.NewOsFs())
	server.setLogger()
	server.migrate()
	return server
}

func (s *Server) setLogger() {
	level := config.GetString("logging.level", "warn")
	log.NewLogger(level).Load()
	log.Info("config.initialised", fmt.Sprintf("the env config has been set to '%s'", s.options.Env))
}

func (s *Server) migrate() {
	repo, err := db.NewDefaultRepository()
	if err != nil {
		log.Error("server.repository.error", err)
	}
	err = repo.Migrate(s.options.Entities)
	if err != nil {
		log.Error("server.repository.error", err)
	}
}

// Listen the server and ensure it's configured
func (s *Server) Listen() {
	s.setMode()
	s.setTimeOuts()
	s.router = s.createRouter()
	port := config.GetString("server.port", "1102")
	err := endless.ListenAndServe(":"+port, s.router)
	if err != nil {
		log.Info("server.shutdown", "server has shutdown.  Goodbye!")
	}
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
	timeOuts := newTimeOuts()
	endless.DefaultReadTimeOut = timeOuts.read
	endless.DefaultWriteTimeOut = timeOuts.write
	endless.DefaultHammerTime = timeOuts.hammer
	log.Info("server.timeouts", "server timeout configuration completed")
}
