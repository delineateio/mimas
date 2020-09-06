package db

import (
	"errors"
	"fmt"

	"github.com/delineateio/mimas/env"
	e "github.com/delineateio/mimas/errors"
)

// DefaultType is the default database type
const DefaultType = "postgres"

// DefaultConnectionString is the default string for formatting
const DefaultConnectionString = "postgres://%s:%s@localhost:5432/%s?sslmode=disable"

// Info represents the connection details
type Info struct {
	Name     string
	Type     string
	Username string
	Password string
}

// NewDefaultInfo creates database info type
func NewDefaultInfo() (*Info, error) {
	vars := env.NewEnv()
	name, err := vars.ReadRequired("DB_NAME")
	if err != nil {
		return nil, err
	}
	username, err := vars.ReadRequired("DB_USERNAME")
	if err != nil {
		return nil, err
	}
	password, err := vars.ReadRequired("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	return NewInfo(name, username, password)
}

// NewInfo creates database info type
func NewInfo(name, username, password string) (*Info, error) {
	errs := e.NewErrors()
	info := &Info{
		Type:     DefaultType,
		Name:     readParam("name", name, errs),
		Username: readParam("username", username, errs),
		Password: readParam("password", password, errs),
	}
	if errs.HasErrors() {
		return nil, errors.New("errors configuring the db")
	}
	return info, nil
}

func readParam(name, value string, errs *e.Errors) string {
	if value == "" {
		errs.Add("db.info.error."+name, fmt.Errorf("parameter '%s' not provided", name))
	}
	return value
}

// ConnectionString returns the formatted connection string
func (info *Info) ConnectionString() string {
	return fmt.Sprintf(DefaultConnectionString, info.Username, info.Password, info.Name)
}
