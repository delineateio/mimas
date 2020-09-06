package config

import (
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var file = `string:
  found: "yes"
bool:
  found: true
  parse:
    error: "efk"
int:
  found: 2
  parse:
    error: "xyx"
duration:
  found: 1s
  parse:
    error: "abc"
uint:
  found: 2
  parse:
    error: -1
strings:
  - first
  - second`

func loadUnitTestConfig(t *testing.T) {
	var configurator = Configurator{
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

func TestUTConfigNotFound(t *testing.T) {
	c := NewConfigurator("", "", afero.NewMemMapFs())
	assert.Panics(t, func() {
		c.Load()
	})
}

func TestUTStringDefaultWithoutLoad(t *testing.T) {
	value := GetString("string.default", "test")
	assert.Equal(t, viper.InConfig("key"), false)
	assert.Equal(t, value, "test")
}

func TestUTStringDefaultWithLoad(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetString("string.default", "test")
	assert.Equal(t, viper.InConfig("key"), false)
	assert.Equal(t, value, "test")
}

func TestUTStringFound(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetString("string.found", "test")
	assert.Equal(t, value, "yes")
	assert.Equal(t, Exists("string.found"), true)
}

func TestUTBoolDefault(t *testing.T) {
	value := GetBool("bool.default", true)
	assert.Equal(t, value, true)
	assert.Equal(t, Exists("bool.default"), false)
}

func TestUTBoolFound(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetBool("bool.found", false)
	assert.Equal(t, value, true)
	assert.Equal(t, Exists("bool.found"), true)
}

func TestUTBoolParseError(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetBool("bool.parse.error", true)
	assert.Equal(t, value, true)
	assert.Equal(t, Exists("bool.parse.error"), true)
}

func TestUTIntDefault(t *testing.T) {
	value := GetInt("int.default", 1)
	assert.Equal(t, value, 1)
	assert.Equal(t, Exists("string.default"), false)
}

func TestUTIntFound(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetInt("int.found", 1)
	assert.Equal(t, value, 2)
	assert.Equal(t, Exists("int.found"), true)
}

func TestUTIntParseError(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetInt("int.parse.error", 1)
	assert.Equal(t, value, 1)
	assert.Equal(t, Exists("int.parse.error"), true)
}

func TestUTDurationDefault(t *testing.T) {
	duration := 500 * time.Millisecond
	value := GetDuration("duration.default", duration)
	assert.Equal(t, value, duration)
	assert.Equal(t, Exists("duration.default"), false)
}

func TestUTDurationFound(t *testing.T) {
	duration := 500 * time.Millisecond
	loadUnitTestConfig(t)
	value := GetDuration("duration.found", duration)
	assert.Equal(t, value, time.Second)
	assert.Equal(t, Exists("duration.found"), true)
}

func TestUTDurationParseError(t *testing.T) {
	duration := 500 * time.Millisecond
	loadUnitTestConfig(t)
	value := GetDuration("duration.parse.error", duration)
	assert.Equal(t, value, duration)
	assert.Equal(t, Exists("duration.parse.error"), true)
}

func TestUTUintDefault(t *testing.T) {
	value := GetUint("uint.default", uint(1))
	assert.Equal(t, value, uint(1))
	assert.Equal(t, Exists("uint.default"), false)
}

func TestUTUintFound(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetUint("uint.found", uint(1))
	assert.Equal(t, value, uint(2))
	assert.Equal(t, Exists("uint.found"), true)
}

func TestUTUintParseError(t *testing.T) {
	loadUnitTestConfig(t)
	value := GetUint("uint.parse.error", uint(1))
	assert.Equal(t, value, uint(1))
	assert.Equal(t, Exists("uint.parse.error"), true)
}

func TestUTGetStrings(t *testing.T) {
	loadUnitTestConfig(t)
	values := GetStrings("strings", []string{"*"})
	assert.Equal(t, len(values), 2)
	assert.Equal(t, values[0], "first")
	assert.Equal(t, values[1], "second")
}

func TestUTGetStringsMissing(t *testing.T) {
	loadUnitTestConfig(t)
	values := GetStrings("strings_missing", []string{"*"})
	assert.Equal(t, len(values), 1)
}
