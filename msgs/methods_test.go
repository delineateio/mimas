package msgs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidMethods(t *testing.T) {
	methods := GetValidMethods()
	assert.Equal(t, len(methods), 5)
	assert.Equal(t, methods[0], http.MethodGet)
	assert.Equal(t, methods[1], http.MethodPost)
	assert.Equal(t, methods[2], http.MethodPut)
	assert.Equal(t, methods[3], http.MethodPatch)
	assert.Equal(t, methods[4], http.MethodDelete)
}

func TestAllMethods(t *testing.T) {
	for _, method := range GetValidMethods() {
		req, err := NewRequest(method, nil)
		if assert.NotNil(t, req) {
			assert.Equal(t, method, req.Method)
		}
		assert.Nil(t, err)
	}
}

func TestMethodNotProvided(t *testing.T) {
	err := ValidateMethod("")
	if err != nil {
		assert.NotNil(t, err)
	}
}

func TestInvalidMethod(t *testing.T) {
	err := ValidateMethod(http.MethodOptions)
	if err != nil {
		assert.NotNil(t, err)
	}
}
