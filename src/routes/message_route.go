package routes

import (
	"github.com/AlfaSakan/my-profile-api.git/src/handlers"
	"github.com/AlfaSakan/my-profile-api.git/src/middlewares"

	"github.com/gin-gonic/gin"
)

const MESSAGE_ROUTE = "/message"

func RouteMessage(router *gin.RouterGroup, handler *handlers.MessageHandler) {
	router.GET(MESSAGE_ROUTE+"/:chatRoomId", middlewares.RequireUser(), handler.GetMessageHandler)

	router.POST(MESSAGE_ROUTE, middlewares.RequireUser(), handler.PostMessageHandler)

	router.PATCH(MESSAGE_ROUTE+"/:messageId", handler.PatchMessageHandler)
}
