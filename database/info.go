package database

import (
	"errors"
	"fmt"
	"os"

	log "github.com/delineateio/mimas/log"

	// Used internally by gorm to load the postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// DefaultType is the default database type
const defaultType = "postgres"
const connectionString = "postgres://%s:%s@localhost:5432/%s?sslmode=disable"

// Info represents the connection details
type Info struct {
	Name     string
	Type     string
	Username string
	Password string
}

// NewInfo creates database info type
func NewInfo() *Info {
	return &Info{
		Name:     os.Getenv("DB_NAME"),
		Type:     defaultType,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

// ConnectionString returns the formatted connection string
func (info *Info) ConnectionString() string {
	// Retrieves from env variables
	var err error

	if info.Name == "" || info.Username == "" || info.Password == "" {
		err = errors.New("no connection string was provided")
		log.Error("db.connection.error", err)
	}

	return fmt.Sprintf(connectionString, info.Username, info.Password, info.Name)
}
