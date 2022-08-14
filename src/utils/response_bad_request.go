package utils

import (
	"myProfileApi/src/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(ctx *gin.Context, response *schemas.Response, err error) {
	response.Message = err.Error()
	response.Status = http.StatusBadRequest
	ctx.JSON(response.Status, response)
}
