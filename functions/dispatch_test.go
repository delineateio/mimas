package functions

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/delineateio/mimas/handlers"
	"github.com/delineateio/mimas/messages"
	"github.com/stretchr/testify/assert"
)

type ErrorReader int

func (ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

type TestObj struct {
	Message string `json:"message" binding:"required"`
}

// NonJSONHander is a handler that can be used for testing purposes
func NonJSONHandler(request *messages.Request, response *messages.Response) {
	response.Body = "hello,world!"
	response.Code = http.StatusOK
}

func JSONHandler(request *messages.Request, response *messages.Response) {
	response.Body = TestObj{Message: "hello,world!"}
	response.Code = http.StatusOK
}

func TestSuccessfulHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatch(w, r, handlers.HealthzHandler)
	})
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, 0, recorder.Body.Len())
}

func TestFailedReaderError(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", new(ErrorReader))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatch(w, r, handlers.NullHandler)
	})
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestNonJSONResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatch(w, r, NonJSONHandler)
	})
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func TestJSONResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatch(w, r, JSONHandler)
	})
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
