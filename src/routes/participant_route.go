package routes

import (
	"fmt"

	"github.com/AlfaSakan/my-profile-api.git/src/handlers"
	"github.com/AlfaSakan/my-profile-api.git/src/middlewares"

	"github.com/gin-gonic/gin"
)

const PARTICIPANT_ROUTE = "/participant"

func ParticipantRoute(router *gin.RouterGroup, chatRoomHandler *handlers.ChatRoomHandler) {
	router.GET(fmt.Sprintf("%s/%s", PARTICIPANT_ROUTE, ":chatRoomId"), middlewares.RequireUser(), chatRoomHandler.GetParticipantsInChatRoom) // DONE
}
