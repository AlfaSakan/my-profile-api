package routes

import (
	"fmt"

	"github.com/AlfaSakan/my-profile-api.git/src/handlers"
	"github.com/AlfaSakan/my-profile-api.git/src/middlewares"

	"github.com/gin-gonic/gin"
)

const USER_ROUTE = "/user"

func RouteUser(router *gin.RouterGroup, handler *handlers.UserHandler) {
	router.GET(USER_ROUTE, middlewares.RequireUser(), handler.GetUserHandler)

	router.GET(fmt.Sprintf("%s/find/:UserId", USER_ROUTE), handler.GetOneUserHandler)

	router.GET(fmt.Sprintf("%s/:Name", USER_ROUTE), middlewares.RequireUser(), handler.SearchUserHandler)

	router.POST(USER_ROUTE, handler.PostUserHandler)

	router.PATCH(USER_ROUTE, handler.PatchUserHandler)
}
