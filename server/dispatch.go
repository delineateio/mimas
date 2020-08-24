package server

import (
	"net/http"

	log "github.com/delineateio/mimas/log"
	messages "github.com/delineateio/mimas/messages"
	"github.com/gin-gonic/gin"
)

// Dispatch initiates and handles the request and response
func Dispatch(ctx *gin.Context, command messages.Command) {
	request := messages.Request{
		Body: make(map[string]interface{}),
	}

	err := ctx.ShouldBind(&request.Body)
	if err != nil {
		log.Error("request.bind.error", err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	var response messages.Response
	command(&request, &response)

	// Adds the headers
	ctx.Header("Content-Type", "application/json")

	// Only writes the body
	if response.Body == nil {
		ctx.Writer.WriteHeader(response.Code)
	} else {
		ctx.JSON(response.Code, response.Body)
	}
}
