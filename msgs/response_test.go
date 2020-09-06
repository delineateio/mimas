package msgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestObj struct {
	Message string `json:"message" binding:"required"`
}

func TestJSONReponse(t *testing.T) {
	response := NewJSONResponse()
	assert.Equal(t, len(response.Headers), 1)
	assert.Equal(t, response.Code, 0)
	assert.Equal(t, response.Headers["Content-Type"], "application/json")
	assert.False(t, response.HasBody())
}

func TestJSONResponseNoHeader(t *testing.T) {
	response := NewJSONResponse()
	assert.Equal(t, len(response.Headers), 1)
	assert.Equal(t, response.Headers["token"], "")
}

func TestJSONResponseHasBodyFalse(t *testing.T) {
	response := NewJSONResponse()
	assert.Nil(t, response.Body)
	assert.False(t, response.HasBody())
}

func TestJSONResponseNonJsonBody(t *testing.T) {
	response := NewJSONResponse()
	response.Body = "hello,world!"
	assert.True(t, response.HasBody())
	result := response.IsValid()
	assert.False(t, result)
}

func TestJSONResponseJsonBody(t *testing.T) {
	response := NewJSONResponse()
	response.Body = TestObj{Message: "hello,world!"}
	assert.True(t, response.HasBody())
	result := response.IsValid()
	assert.True(t, result)
}

func TestJSONResponseToBytes(t *testing.T) {
	response := NewJSONResponse()
	result := response.ToBytes()
	assert.Nil(t, result)
}
