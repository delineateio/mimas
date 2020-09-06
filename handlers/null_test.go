package handlers

import (
	"net/http"
	"testing"

	"github.com/delineateio/mimas/msgs"
	"github.com/stretchr/testify/assert"
)

func TestNullMissingRequest(t *testing.T) {
	res := &msgs.Response{}
	NullHandler(nil, res)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestNullSuccessfulResponse(t *testing.T) {
	res := &msgs.Response{}
	NullHandler(&msgs.Request{}, res)
	assert.Equal(t, http.StatusOK, res.Code)
}
