package cors

import (
	"net/http"
	"testing"
	"time"

	"github.com/delineateio/mimas/config"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

// Notes spaces are needed NOT tabs
var file = `server:
  cors:
    allow_origins:
      - https://www.delineate.dev
      - https://www.delineate.pub
      - https://www.delineate.io
    allow_methods:
      - GET
      - POST
    allow_headers:
      - Origin
    expose_headers:
      - Content-Length
    allow_credentials: true
    max_age: 12h`

func loadUnitTestConfig(t *testing.T) {
	var configurator = config.Configurator{
		Env:      "config",
		Location: "/",
		Fs:       getFs(t),
	}
	configurator.Load()
}

func getFs(t *testing.T) afero.Fs {
	fs := afero.NewMemMapFs()
	err := fs.MkdirAll("/", 0755)
	assert.Nil(t, err)
	err = afero.WriteFile(fs, "/config.yml", []byte(file), 0644)
	assert.Nil(t, err)
	return fs
}

func TestCORSAllowOrigins(t *testing.T) {
	loadUnitTestConfig(t)
	corsConfig := NewCORSConfig()
	allowOrigins := corsConfig.AllowOrigins
	assert.Equal(t, 3, len(allowOrigins))
	assert.Equal(t, "https://www.delineate.dev", allowOrigins[0])
	assert.Equal(t, "https://www.delineate.pub", allowOrigins[1])
	assert.Equal(t, "https://www.delineate.io", allowOrigins[2])
}

func TestCORSAllowMethods(t *testing.T) {
	loadUnitTestConfig(t)
	corsConfig := NewCORSConfig()
	allowMethods := corsConfig.AllowMethods
	assert.Equal(t, 2, len(allowMethods))
	assert.Equal(t, http.MethodGet, allowMethods[0])
	assert.Equal(t, http.MethodPost, allowMethods[1])
}

func TestCORSAllowHeaders(t *testing.T) {
	loadUnitTestConfig(t)
	corsConfig := NewCORSConfig()
	allowHeaders := corsConfig.AllowHeaders
	assert.Equal(t, 1, len(allowHeaders))
	assert.Equal(t, "Origin", allowHeaders[0])
}

func TestCORSExposeHeaders(t *testing.T) {
	loadUnitTestConfig(t)
	corsConfig := NewCORSConfig()
	exposeHeaders := corsConfig.ExposeHeaders
	assert.Equal(t, 1, len(exposeHeaders))
	assert.Equal(t, "Content-Length", exposeHeaders[0])
}

func TestCORSAllowCredentials(t *testing.T) {
	loadUnitTestConfig(t)
	corsConfig := NewCORSConfig()
	allow := corsConfig.AllowCredentials
	assert.True(t, allow)
}

func TestCORSMaxAge(t *testing.T) {
	loadUnitTestConfig(t)
	corsConfig := NewCORSConfig()
	maxAge := corsConfig.MaxAge
	assert.Equal(t, 12*time.Hour, maxAge)
}
