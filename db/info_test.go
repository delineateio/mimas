package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const dbtype = "type"
const name = "name"
const username = "username"
const password = "pw"

func getInfo() *Info {
	return &Info{
		Type:     dbtype,
		Name:     name,
		Username: username,
		Password: password,
	}
}

func TestInfo(t *testing.T) {
	info := getInfo()
	assert.Equal(t, dbtype, info.Type)
	assert.Equal(t, name, info.Name)
	assert.Equal(t, username, info.Username)
	assert.Equal(t, password, info.Password)
}

func TestInfoConnectionString(t *testing.T) {
	info := getInfo()
	cs := info.ConnectionString()
	expected := fmt.Sprintf(DefaultConnectionString, info.Username, info.Password, info.Name)
	assert.Equal(t, expected, cs)
}

func TestInfoParamMissing(t *testing.T) {
	info, err := NewInfo("", "", "")
	assert.NotNil(t, err)
	assert.Nil(t, info)
}

func TestInfoFromParams(t *testing.T) {
	info, err := NewInfo("postgres", "postgres", "postgres")
	assert.NotNil(t, info)
	assert.Nil(t, err)
}
