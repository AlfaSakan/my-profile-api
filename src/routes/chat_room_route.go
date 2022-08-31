package routes

import (
	"fmt"

	"github.com/AlfaSakan/my-profile-api.git/src/handlers"
	"github.com/AlfaSakan/my-profile-api.git/src/middlewares"
	"github.com/AlfaSakan/my-profile-api.git/src/websocket"

	"github.com/gin-gonic/gin"
)

const CHAT_ROOM_ROUTE = "/chat-room"

func ChatRoomRoute(router *gin.RouterGroup, chatRoomHandler *handlers.ChatRoomHandler, h *websocket.Hub) {
	router.GET(fmt.Sprintf("/ws%s/:UserId", CHAT_ROOM_ROUTE), chatRoomHandler.WebSocketGetAllChatRoom(h)) // DONE

	router.GET(CHAT_ROOM_ROUTE, middlewares.RequireUser(), chatRoomHandler.GetAllChatRoom) // DONE

	router.POST(CHAT_ROOM_ROUTE, middlewares.RequireUser(), chatRoomHandler.PostChatRoom(h)) // DONE

	router.POST(fmt.Sprintf("%s/%s", CHAT_ROOM_ROUTE, "add-participant"), middlewares.RequireUser(), chatRoomHandler.PostAddParticipant(h)) // DONE

	router.DELETE(fmt.Sprintf("%s/%s", CHAT_ROOM_ROUTE, ":chatRoomId"), middlewares.RequireUser(), chatRoomHandler.DeleteChatRoom) // DONE

	router.PATCH(CHAT_ROOM_ROUTE+"/:chatRoomId", chatRoomHandler.PatchChatRoom)

	router.GET(fmt.Sprintf("%s/%s", CHAT_ROOM_ROUTE, ":chatRoomId"), middlewares.RequireUser(), chatRoomHandler.GetChatRoomById) // DONE
}
