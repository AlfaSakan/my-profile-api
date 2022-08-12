package routes

import (
	"myProfileApi/src/handlers"

	"github.com/gin-gonic/gin"
)

const USER_ROUTE = "/user"

func RouteUser(router *gin.RouterGroup, handler *handlers.UserHandler) {
	router.GET(USER_ROUTE+"/:userId", handler.GetUserHandler)

	router.POST(USER_ROUTE, handler.PostUserHandler)

	router.PATCH(USER_ROUTE+"/:userId", handler.PatchUserHandler)
}
