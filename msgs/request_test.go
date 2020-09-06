package msgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestMissingMethod(t *testing.T) {
	req, err := NewRequest("", nil)
	assert.Nil(t, req)
	assert.NotNil(t, err)
}

func TestRequestIncorrectMethod(t *testing.T) {
	req, err := NewRequest("FAKE", nil)
	assert.Nil(t, req)
	assert.NotNil(t, err)
}

func TestRequestCorrectMethod(t *testing.T) {
	for _, method := range GetValidMethods() {
		req, err := NewRequest(method, nil)
		if assert.NotNil(t, req) {
			assert.Equal(t, method, req.Method)
		}
		assert.Nil(t, err)
	}
}
