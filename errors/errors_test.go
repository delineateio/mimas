package errors

import (
	e "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const event = "test.error"
const desc = "new error"

func TestEmptyErrors(t *testing.T) {
	errors := NewErrors()
	assert.NotNil(t, errors)
	assert.Equal(t, len(errors.Items), 0)
	assert.False(t, errors.HasErrors())
}

func TestAddNilError(t *testing.T) {
	errors := NewErrors()
	errors.Add("key", nil)
	assert.NotNil(t, errors)
	assert.Equal(t, len(errors.Items), 0)
	assert.False(t, errors.HasErrors())
}

func TestAddError(t *testing.T) {
	errors := NewErrors()
	assert.NotNil(t, errors)
	errors.Add(event, e.New(desc))
	assert.Equal(t, len(errors.Items), 1)
	assert.Equal(t, errors.Items[0].Error(), desc)
	assert.True(t, errors.HasErrors())
}

func TestAddMultipleErrors(t *testing.T) {
	errors := NewErrors()
	assert.NotNil(t, errors)
	errors.Add(desc, e.New(desc))
	errors.Add(desc, e.New(desc))
	errors.Add(desc, e.New(desc))
	assert.Equal(t, len(errors.Items), 3)
	assert.True(t, errors.HasErrors())
}

func TestAddErrorWithoutEvent(t *testing.T) {
	errors := NewErrors()
	assert.NotNil(t, errors)
	errors.Add("", e.New(desc))
	assert.Equal(t, len(errors.Items), 1)
	assert.Equal(t, errors.Items[0].Error(), desc)
	assert.True(t, errors.HasErrors())
}
