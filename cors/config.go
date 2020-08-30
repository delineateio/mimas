package cors

import (
	"time"

	"github.com/delineateio/mimas/config"
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
		AllowOrigins:     config.GetStrings("server.cors.allow_origins"),
		AllowMethods:     config.GetStrings("server.cors.allow_methods"),
		AllowHeaders:     config.GetStrings("server.cors.allow_headers"),
		ExposeHeaders:    config.GetStrings("server.cors.expose_headers"),
		AllowCredentials: config.GetBool("server.cors.allow_credentials", false),
		MaxAge:           config.GetDuration("server.cors.allow_credentials", time.Hour),
	}
}
