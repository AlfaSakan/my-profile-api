package routes

import (
	"fmt"
	"myProfileApi/src/handlers"
	"myProfileApi/src/middlewares"

	"github.com/gin-gonic/gin"
)

const CHAT_ROOM_ROUTE = "/chat-room"

func ChatRoomRoute(router *gin.RouterGroup, chatRoomHandler *handlers.ChatRoomHandler) {
	router.GET(CHAT_ROOM_ROUTE, middlewares.RequireUser(), chatRoomHandler.GetAllChatRoom) // DONE

	router.POST(CHAT_ROOM_ROUTE, middlewares.RequireUser(), chatRoomHandler.PostChatRoom) // DONE

	router.POST(fmt.Sprintf("%s/%s", CHAT_ROOM_ROUTE, "add-participant"), middlewares.RequireUser(), chatRoomHandler.PostAddParticipant) // DONE

	router.DELETE(fmt.Sprintf("%s/%s", CHAT_ROOM_ROUTE, ":chatRoomId"), middlewares.RequireUser(), chatRoomHandler.DeleteChatRoom) // DONE

	router.PATCH(CHAT_ROOM_ROUTE+"/:chatRoomId", chatRoomHandler.PatchChatRoom)

	router.GET(fmt.Sprintf("%s/%s", CHAT_ROOM_ROUTE, ":chatRoomId"), middlewares.RequireUser(), chatRoomHandler.GetChatRoomById) // DONE
}
