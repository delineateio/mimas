package cors

import (
	"time"

	"github.com/delineateio/mimas/config"
	"github.com/delineateio/mimas/msgs"
)

// Config provides the CORS config
type Config struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           time.Duration
}

// NewCORSConfig loads the config from the configuration store
func NewCORSConfig() *Config {
	return &Config{
		AllowOrigins:     config.GetStrings("server.cors.allow_origins", []string{"*"}),
		AllowMethods:     config.GetStrings("server.cors.allow_methods", msgs.GetValidMethods()),
		AllowHeaders:     config.GetStrings("server.cors.allow_headers", []string{}),
		ExposeHeaders:    config.GetStrings("server.cors.expose_headers", []string{}),
		AllowCredentials: config.GetBool("server.cors.allow_credentials", false),
		MaxAge:           config.GetDuration("server.cors.max_age", time.Hour),
	}
}
