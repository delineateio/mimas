package handlers

import (
	"net/http"

	"github.com/delineateio/mimas/db"
	e "github.com/delineateio/mimas/errors"
	"github.com/delineateio/mimas/msgs"
)

// Handler performs the required action for the service
type Handler func(request *msgs.Request, response *msgs.Response)

// Create performs a simple CRUD operation into the DB
func Create(request *msgs.Request, entity interface{}, response *msgs.Response) {
	// Translates the request
	errs := e.NewErrors()
	err := request.Translate(&entity)
	checkError("request.error", err, errs, response, http.StatusBadRequest)
	// Creates the repo
	repo, err := db.NewDefaultRepository()
	checkError("db.error", err, errs, response, http.StatusServiceUnavailable)
	// Creates the entity
	err = repo.Create(entity)
	checkError("db.error.create", err, errs, response, http.StatusInternalServerError)
	// If no errors returns
	if !errs.HasErrors() {
		response.Code = http.StatusCreated
	}
}

func checkError(event string, err error, errs *e.Errors, response *msgs.Response, code int) {
	if err != nil {
		errs.Add(event, err)
		response.Code = code
	}
}
