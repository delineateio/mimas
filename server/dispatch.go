package server

import (
	"net/http"

	"github.com/delineateio/mimas/errors"
	"github.com/delineateio/mimas/handlers"
	"github.com/delineateio/mimas/msgs"
	"github.com/gin-gonic/gin"
)

func dispatch(ctx *gin.Context, handler handlers.Handler) {
	// Creates the request
	errs := errors.NewErrors()
	request, err := msgs.NewRequest(ctx.Request.Method, ctx.Request.Header)
	errs.Add("request.create.error", err)

	err = ctx.ShouldBind(&request.Body)
	errs.Add("request.bind.error", err)

	if errs.HasErrors() {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
	} else {
		response := msgs.NewJSONResponse()
		handler(request, response)
		setResponse(ctx, response)
	}
}

func setResponse(ctx *gin.Context, response *msgs.Response) {
	for key, value := range response.Headers {
		ctx.Header(key, value)
	}
	if response.Body == nil {
		ctx.Writer.WriteHeader(response.Code)
	} else {
		ctx.JSON(response.Code, response.Body)
	}
}
