package messages

import (
	"errors"
	"fmt"
	"net/http"
)

// ValidateMethod validates that the methods is in the the right set
func ValidateMethod(method string) error {
	if method == "" {
		return errors.New("no HTTP method was provided")
	}
	switch method {
	case
		http.MethodGet,
		http.MethodPost,
		http.MethodDelete,
		http.MethodPut,
		http.MethodPatch:
		return nil
	}
	return fmt.Errorf("'%s' is not a valid HTTP method", method)
}

// GetValidMethods returns allowed superset of the HTTP methods
func GetValidMethods() []string {
	return []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}
}
