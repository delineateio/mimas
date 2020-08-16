package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const Default = "postgres"
const None = "none"

func loadConfigUnitTestConfig() {
	var configurator = Configurator{
		Env:      "repository",
		Location: "../config",
	}
	configurator.Load()
}

func TestUTDefaultDatabaseType(t *testing.T) {
	r := NewRepository("missing")
	dbType, err := r.getDatabaseType()

	assert.Equal(t, dbType, DefaultDBType)
	assert.NoError(t, err)
}

func TestUTGetConfigDatabaseType(t *testing.T) {
	loadConfigUnitTestConfig()
	r := NewRepository("mysql")
	r.AllowedDBTypes = []string{DefaultDBType, "mysql"}
	dbType, err := r.getDatabaseType()

	assert.NoError(t, err)
	assert.Equal(t, dbType, "mysql")
}

func TestUTDBTypeNotAllowed(t *testing.T) {
	loadConfigUnitTestConfig()
	r := NewRepository("sqllite")
	r.AllowedDBTypes = []string{DefaultDBType, "mysql"}
	dbType, err := r.getDatabaseType()

	assert.Error(t, err)
	assert.Equal(t, dbType, DefaultDBType)
}

func TestUTGetConfigConnectionString(t *testing.T) {
	loadConfigUnitTestConfig()
	r := NewRepository("postgres")
	r.Username = Default
	// pragma: allowlist secret
	r.Password = Default
	r.DBName = Default
	connection, err := r.getConnectionString()

	assert.NoError(t, err)
	assert.Equal(t, connection, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
}

func TestUTOpenTryAttempts(t *testing.T) {
	loadConfigUnitTestConfig()
	r := NewRepository("postgres")
	r.Username = None
	r.Password = None
	r.DBName = None
	assert.Equal(t, r.Attempts, uint(3))
	// Changes from the default
	attempts := 2
	r.Attempts = uint(attempts)

	assert.Error(t, r.Open())
	assert.Equal(t, r.Info.Tries, attempts)
}
