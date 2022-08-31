package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/services"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"
	"github.com/AlfaSakan/my-profile-api.git/src/websocket"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ChatRoomHandler struct {
	chatRoomService    services.IChatRoomService
	participantService services.IParticipantService
}

func NewChatRoomHandler(chatRoomService services.IChatRoomService, participantService services.IParticipantService) *ChatRoomHandler {
	return &ChatRoomHandler{chatRoomService, participantService}
}

func (chatRoomHandler *ChatRoomHandler) GetAllChatRoom(ctx *gin.Context) {
	response := new(schemas.Response)

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	chatRooms, err := chatRoomHandler.chatRoomService.FindAllChatRoomByUserId(userId)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusBadRequest

		ctx.JSON(response.Status, response)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = &chatRooms

	ctx.JSON(response.Status, response)
}

func (chatRoomHandler *ChatRoomHandler) WebSocketGetAllChatRoom(h *websocket.Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("UserId")

		chatRooms, _ := chatRoomHandler.chatRoomService.FindAllChatRoomByUserId(userId)

		if len(chatRooms) == 0 {
			schema := schemas.ChatRoomWithPartisipants{
				ChatRoomId:     "0",
				ParticipantsId: &[]string{"0"},
				Messages:       &[]models.Message{},
				ImageUrl:       "",
				Description:    "",
				Name:           "",
				Type:           "noreply",
				CreatedAt:      time.Now().UnixMilli(),
				UpdatedAt:      time.Now().UnixMilli(),
			}

			chatRooms = append(chatRooms, schema)
		}

		websocket.ServeWs(ctx, h, &chatRooms)
	}
}

func (chatRoomHandler *ChatRoomHandler) PostChatRoom(h *websocket.Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request schemas.ChatRoomRequest
		response := new(schemas.Response)

		user, _ := ctx.Get("User")
		userId := user.(*models.User).UserId

		request.UserId = userId

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				response.Message = e.Error()
				response.Status = http.StatusBadRequest

				ctx.JSON(response.Status, response)
				return
			}
		}

		chatRoom, errChatRoom := chatRoomHandler.chatRoomService.CreateChatRoom(request)
		if errChatRoom != nil {
			response.Message = errChatRoom.Error()
			response.Status = http.StatusBadRequest

			ctx.JSON(response.Status, response)
			return
		}

		addParticipantRequest := &schemas.AddParticipantRequest{
			ChatRoomId: chatRoom.ChatRoomId,
			UserIds:    *chatRoom.ParticipantsId,
		}

		messageModel := &models.Message{
			ChatRoomId: chatRoom.ChatRoomId,
			Type:       "Created Room",
		}

		message, _ := json.Marshal(messageModel)

		data := map[string][]byte{
			"message": message,
			"id":      []byte(userId),
		}

		userRoom, _ := json.Marshal(data)

		h.AddChatRoom <- addParticipantRequest
		h.Broadcast <- userRoom

		response.Data = &chatRoom
		response.Message = "Created"
		response.Status = http.StatusCreated
		ctx.JSON(response.Status, response)
	}
}

func (chatRoomHandler *ChatRoomHandler) PatchChatRoom(ctx *gin.Context) {
	var request schemas.ChatRoomRequest
	response := new(schemas.Response)

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	request.UserId = userId

	chatRoomId := ctx.Param("chatRoomId")

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			response.Message = e.Error()
			response.Status = http.StatusBadRequest

			ctx.JSON(response.Status, response)
			return
		}
	}

	errHandler := chatRoomHandler.chatRoomService.UpdateChatRoom(&request, chatRoomId)
	if errHandler != nil {
		response.Message = errHandler.Error()
		response.Status = http.StatusBadRequest

		ctx.JSON(response.Status, response)
		return
	}

	response.Data = ""
	response.Status = http.StatusOK
	response.Message = "Success updated"

	ctx.JSON(response.Status, response)
}

func (chatRoomHandler *ChatRoomHandler) DeleteChatRoom(ctx *gin.Context) {
	response := new(schemas.Response)
	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	chatRoomId := ctx.Param("chatRoomId")

	err := chatRoomHandler.chatRoomService.RemoveChatRoom(userId, chatRoomId)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Status = http.StatusOK
	response.Data = ""
	response.Message = "Success deleted"

	ctx.JSON(response.Status, response)
}

func (chatRoomHandler *ChatRoomHandler) PostAddParticipant(h *websocket.Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &schemas.AddParticipantRequest{}
		response := &schemas.Response{}

		err := ctx.ShouldBindJSON(request)
		if err != nil {
			utils.ResponseBadRequest(ctx, response, err)
			return
		}

		user, _ := ctx.Get("User")
		userId := user.(*models.User).UserId

		participant, err := chatRoomHandler.participantService.FindUserAdmin(userId, request.ChatRoomId)

		if participant == nil {
			utils.ResponseBadRequest(ctx, response, err)
			return
		}

		for _, userId := range request.UserIds {
			chatRoomHandler.participantService.CreateParticipant(userId, request.ChatRoomId)
		}

		h.AddChatRoom <- request

		response.Message = "OK"
		response.Status = http.StatusOK
		response.Data = ""
		ctx.JSON(response.Status, response)
	}
}

func (chatRoomHandler *ChatRoomHandler) GetChatRoomById(ctx *gin.Context) {
	response := &schemas.Response{}
	chatRoomId := ctx.Param("chatRoomId")

	chatRoom, err := chatRoomHandler.chatRoomService.FindChatRoomById(chatRoomId)

	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = chatRoom
	ctx.JSON(response.Status, response)
}

func (chatRoomHandler *ChatRoomHandler) GetParticipantsInChatRoom(ctx *gin.Context) {
	response := &schemas.Response{}
	chatRoomId := ctx.Param("chatRoomId")

	users, err := chatRoomHandler.participantService.FindParticipantsList(chatRoomId)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = users
	ctx.JSON(response.Status, response)
}
