package environment

import (
	"fmt"
	"os"

	"github.com/delineateio/mimas/log"
)

const readEnvEvent = "environment.env.read"
const readEnvDesc = "Env '%s' didn't exist so defaulted to '%s'"

// Env represents the configured
type Env struct{}

// NewEnv creates a new env
func NewEnv() *Env {
	return &Env{}
}

// ReadEnv env variable from
func (e *Env) Read(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Warn(readEnvEvent, fmt.Sprintf(readEnvDesc, key, defaultValue))
		return defaultValue
	}
	return value
}

// ReadRequired gets a required env variable or throws an error
func (e *Env) ReadRequired(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("env '%s' does not exist", key)
	}

	return value, nil
}
