package handlers

import (
	"net/http"
	"testing"

	"github.com/delineateio/mimas/messages"
	"github.com/stretchr/testify/assert"
)

func TestNullMissingRequest(t *testing.T) {
	res := &messages.Response{}
	NullHandler(nil, res)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestNullSuccessfulResponse(t *testing.T) {
	res := &messages.Response{}
	NullHandler(&messages.Request{}, res)
	assert.Equal(t, http.StatusOK, res.Code)
}
