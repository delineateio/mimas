package server

import (
	"net/http"

	c "github.com/delineateio/mimas/common"
	"github.com/gin-gonic/gin"
)

// Dispatch initiates and handles the request and response
func Dispatch(ctx *gin.Context, command c.Command) {
	request := c.Request{
		Body: make(map[string]interface{}),
	}

	err := ctx.ShouldBind(&request.Body)
	if err != nil {
		c.Error("request.bind.error", err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	var response c.Response
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
