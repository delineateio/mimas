package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const key = "TEST_ENV"
const dev = "dev"
const io = "io"

func TestReadEnvDefault(t *testing.T) {
	env := NewEnv()
	value := env.Read(key, dev)
	assert.Equal(t, value, dev)
}

func TestReadEnv(t *testing.T) {
	os.Setenv(key, io)
	env := NewEnv()
	value := env.Read(key, dev)
	assert.Equal(t, io, value)
	os.Unsetenv(key)
}

func TestReadRequiredEnv(t *testing.T) {
	os.Setenv(key, io)
	env := NewEnv()
	value, err := env.ReadRequired(key)
	assert.Nil(t, err)
	assert.Equal(t, io, value)
	os.Unsetenv(value)
}

func TestReadRequiredEnvMissing(t *testing.T) {
	env := NewEnv()
	value, err := env.ReadRequired(key + "_missing")
	assert.NotNil(t, err)
	assert.Equal(t, value, "")
}
