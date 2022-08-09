package routes

import (
	"myProfileApi/src/handlers"

	"github.com/gin-gonic/gin"
)

const MESSAGE_ROUTE = "/message"

func RouteMessage(router *gin.RouterGroup, handler *handlers.MessageHandler) {
	router.GET(MESSAGE_ROUTE+"/:chatRoomId", handler.GetMessageHandler)

	router.POST(MESSAGE_ROUTE, handler.PostMessageHandler)

	router.PATCH(MESSAGE_ROUTE+"/:messageId", handler.PatchMessageHandler)
}
