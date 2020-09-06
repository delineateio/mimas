package db

import (
	"testing"
	"time"

	"github.com/delineateio/mimas/config"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

// Notes spaces are needed NOT tabs
var file = `db:
  current:
    limits:
      maxIdle: 10
      maxOpen: 20
      maxLifetime: 1m
    retries:
      attempts: 10
      delay: 3s`

// "db.current.limits.maxIdle"
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

func TestDBSettingDefaults(t *testing.T) {
	settings := NewSettings("none")
	assert.Equal(t, maxIdle, settings.MaxIdle)
	assert.Equal(t, maxOpen, settings.MaxOpen)
	assert.Equal(t, maxLifetime, settings.MaxLifetime)
	assert.Equal(t, attempts, settings.Attempts)
	assert.Equal(t, delay, settings.Delay)
}

func TestDBSettingConfiguration(t *testing.T) {
	loadUnitTestConfig(t)
	settings := NewSettings("current")
	assert.Equal(t, 10, settings.MaxIdle)
	assert.Equal(t, 20, settings.MaxOpen)
	assert.Equal(t, 1*time.Minute, settings.MaxLifetime)
	assert.Equal(t, uint(10), settings.Attempts)
	assert.Equal(t, 3*time.Second, settings.Delay)
}
