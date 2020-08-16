package common

import (
	"strconv"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// IConfigurator interface for injectiong configuration set up
type IConfigurator interface {
	Load(func(in fsnotify.Event))
}

// NewConfigurator gets a new production configurator
func NewConfigurator(env, location string) *Configurator {
	if env == "" {
		env = "io"
	}

	if location == "" {
		location = "/config"
	}

	var configurator = &Configurator{
		Env:      env,
		Location: location,
	}

	return configurator
}

// Configurator sets up configuration in production
type Configurator struct {
	Env      string
	Location string
}

// Load loads without a callback
func (c *Configurator) Load() {
	c.LoadWithCallback(nil)
}

// LoadWithCallback loads up the configuration from the sources
func (c *Configurator) LoadWithCallback(reload func(in fsnotify.Event)) {
	viper.SetConfigType("yml")
	viper.SetConfigName(c.Env)
	viper.AddConfigPath(c.Location)

	// Adds the func for callback (if provided)
	viper.WatchConfig()
	if reload != nil {
		viper.OnConfigChange(reload)

		// This will use the new log level that has been set
		Info("configuration.reload", "the configiration has been reload")
	}

	// Panics if can't be read correctly
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// GetBool gets the boolean value or defaults as required
func GetBool(key string, defaultValue bool) bool {
	if viper.IsSet(key) {
		raw := viper.GetString(key)
		value, err := strconv.ParseBool(raw)
		if err != nil {
			Error(key, err)
			return defaultValue
		}

		Debug(key, strconv.FormatBool(value))
		return value
	}

	Warn(key, "not found in the configuration file, using default")
	return defaultValue
}

// GetString gets the value from Viper
func GetString(key, defaultValue string) string {
	if viper.IsSet(key) {
		value := viper.GetString(key)
		Debug(key, value)
		return value
	}

	Debug(key, "not found in the configuration file, using default")
	return defaultValue
}

// GetInt gets the value from Viper
func GetInt(key string, defaultNumber int) int {
	if viper.IsSet(key) {
		value := viper.GetString(key)
		number, err := strconv.Atoi(value)
		if err != nil {
			Error(key, err)
			return defaultNumber
		}

		Debug(key, strconv.Itoa(number))
		return number
	}

	Debug(key, "not found in the configuration file, using default")
	return defaultNumber
}

// GetDuration provides additional valiation on top of the standard library
// because Viper returned zero duraction which could cause significant performance issues
func GetDuration(key string, defaultDuration time.Duration) time.Duration {
	if viper.IsSet(key) {
		value := viper.GetString(key)
		duration, err := time.ParseDuration(value)
		if err != nil {
			Error(key, err)
			return defaultDuration
		}

		Debug(key, duration.String())
		return duration
	}

	Debug(key, "not found in the configuration file, using default")
	return defaultDuration
}

// GetUint gets the value from Viper
func GetUint(key string, defaultNumber uint) uint {
	if viper.IsSet(key) {
		value := viper.GetString(key)
		number, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			Error(key, err)
			return defaultNumber
		}

		Debug(key, strconv.Itoa(int(number)))
		return uint(number)
	}

	Debug(key, "not found in the configuration file, using default")
	return defaultNumber
}

// Exists confirms if the the key exists
func Exists(key string) bool {
	return viper.IsSet(key)
}
