package utils

import (
	"net/http"
	"strconv"

	"github.com/AlfaSakan/my-profile-api.git/src/schemas"

	"github.com/gin-gonic/gin"
)

func ConvertParamToInt(ctx *gin.Context, key string) int {
	userId, errConvert := strconv.Atoi(ctx.Param(key))

	if errConvert != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response{
			Message: errConvert.Error(),
			Status:  http.StatusBadRequest,
			Data:    "",
		})
		return 0
	}

	return userId
}
