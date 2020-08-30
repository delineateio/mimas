package containers

import (
	"net/http"

	log "github.com/delineateio/mimas/log"
	messages "github.com/delineateio/mimas/messages"
	"github.com/gin-gonic/gin"
)

func dispatch(ctx *gin.Context, command messages.Command) {
	// Creates the request
	request := messages.NewRequest(ctx.Request.Method, ctx.Request.Header)
	err := ctx.ShouldBind(&request.Body)

	if err != nil {
		log.Error("request.bind.error", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
	} else {
		response := messages.NewJSONResponse()
		command(request, response)
		for key, value := range response.Headers {
			ctx.Header(key, value)
		}
		if response.Body == nil {
			ctx.Writer.WriteHeader(response.Code)
		} else {
			ctx.JSON(response.Code, response.Body)
		}
	}
}
