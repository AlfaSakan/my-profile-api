package routes

import (
	"myProfileApi/src/handlers"
	"myProfileApi/src/middlewares"

	"github.com/gin-gonic/gin"
)

const CHAT_ROOM_ROUTE = "/chat-room"

func ChatRoomRoute(router *gin.RouterGroup, chatRoomHandler *handlers.ChatRoomHandler) {
	router.GET(CHAT_ROOM_ROUTE+"/:userId", middlewares.RequireUser(), chatRoomHandler.GetAllChatRoom)

	router.POST(CHAT_ROOM_ROUTE, chatRoomHandler.PostChatRoom)

	router.DELETE(CHAT_ROOM_ROUTE+"/:chatRoomId", chatRoomHandler.DeleteChatRoom)

	router.PATCH(CHAT_ROOM_ROUTE+"/:chatRoomId", chatRoomHandler.PatchChatRoom)
}
