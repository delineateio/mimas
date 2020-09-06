package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Customer represents a customer within this specific domain
type Text struct {
	gorm.Model
	Msg string `json:"msg" binding:"required"`
}

func TestMissingCredentials(t *testing.T) {
	repo, err := getInvalidRepository()
	assert.Nil(t, repo)
	assert.NotNil(t, err)
}

func TestOpenDB(t *testing.T) {
	repo, err := getValidRepository()
	assert.Nil(t, err)
	err = repo.Open()
	assert.Nil(t, err)
}

func TestFailedOpenDB(t *testing.T) {
	repo, err := getRepository("postgres", "postgres", "password")
	assert.Nil(t, err)
	err = repo.Open()
	assert.NotNil(t, err)
}

func TestSuccessfulPingDB(t *testing.T) {
	repo, err := getValidRepository()
	assert.Nil(t, err)
	err = repo.Ping()
	assert.Nil(t, err)
}

func TestFailedPingDB(t *testing.T) {
	repo, err := getRepository("postgres", "postgres", "password")
	assert.Nil(t, err)
	err = repo.Ping()
	assert.NotNil(t, err)
}

func TestNoConnectionCloseDB(t *testing.T) {
	repo, err := getValidRepository()
	assert.Nil(t, err)
	err = repo.Close()
	assert.Nil(t, err)
}

func TestDoubleCloseDB(t *testing.T) {
	repo, err := getValidRepository()
	assert.Nil(t, err)
	err = repo.Open()
	assert.Nil(t, err)
	err = repo.Close()
	assert.Nil(t, err)
	err = repo.Close()
	assert.Nil(t, err)
}

func TestMigrateDB(t *testing.T) {
	repo, err := getValidRepository()
	assert.Nil(t, err)
	var entities []interface{}
	entities = append(entities, &Text{})
	err = repo.Migrate(entities)
	assert.Nil(t, err)
}

func TestCreateDB(t *testing.T) {
	repo, err := getValidRepository()
	assert.Nil(t, err)
	var entities []interface{}
	entities = append(entities, &Text{})
	err = repo.Migrate(entities)
	assert.Nil(t, err)
	err = repo.Create(&Text{Msg: "hello,world!"})
	assert.Nil(t, err)
}

func getValidRepository() (*Repository, error) {
	return getRepository("postgres", "postgres", "postgres")
}

func getInvalidRepository() (*Repository, error) {
	return getRepository("", "", "")
}

func getRepository(name, username, password string) (*Repository, error) {
	return NewRepository(name, username, password)
}
