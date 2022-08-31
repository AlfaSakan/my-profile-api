package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/services"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MessageHandler struct {
	messageService  services.IMessageService
	chatRoomService services.IChatRoomService
}

func NewMessageHandler(messageService services.IMessageService, chatRoomService services.IChatRoomService) *MessageHandler {
	return &MessageHandler{messageService, chatRoomService}
}

func (messageHandler *MessageHandler) GetMessageHandler(ctx *gin.Context) {
	chatRoomId := ctx.Param("chatRoomId")
	response := new(schemas.Response)

	messages, err := messageHandler.messageService.FindMessageByChatRoomId(chatRoomId)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = messages
	response.Status = http.StatusOK
	response.Message = "OK"

	ctx.JSON(http.StatusOK, response)
}

func (messageHandler *MessageHandler) PostMessageHandler(ctx *gin.Context) {
	var request schemas.MessageRequest
	response := new(schemas.Response)

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	request.SenderId = userId

	errRequest := ctx.ShouldBindJSON(&request)
	if errRequest != nil {
		for _, e := range errRequest.(validator.ValidationErrors) {
			response.Status = http.StatusBadRequest
			response.Message = e.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
	}

	participantsId, _ := messageHandler.chatRoomService.FindAllParticipantByChatRoomId(request.ChatRoomId)
	isExist := utils.ArrayContainsUint(participantsId, userId)

	if !isExist {
		utils.ResponseBadRequest(ctx, response, fmt.Errorf("user not in the chat room"))
		return
	}

	responseService, err := messageHandler.messageService.CreateMessage(request)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	chatRoom := &schemas.ChatRoomRequest{UpdatedAt: time.Now().UnixMilli()}

	messageHandler.chatRoomService.UpdateChatRoom(chatRoom, request.ChatRoomId)

	response.Status = http.StatusCreated
	response.Data = responseService
	response.Message = "OK"
	ctx.JSON(http.StatusCreated, response)
}

func (messageHandler *MessageHandler) PatchMessageHandler(ctx *gin.Context) {
	messageId := ctx.Param("messageId")
	response := &schemas.Response{}

	message := &schemas.MessageUpdate{}

	errBindings := ctx.ShouldBindJSON(message)
	if errBindings != nil {
		for _, errBinding := range errBindings.(validator.ValidationErrors) {
			response.Status = http.StatusBadRequest
			response.Message = errBinding.Error()
			ctx.JSON(response.Status, response)
			return
		}
	}

	err := messageHandler.messageService.UpdateMessage(message, messageId)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		ctx.JSON(response.Status, response)
		return
	}

	response.Data = "Success Updated"
	response.Message = "OK"
	response.Status = http.StatusOK
	ctx.JSON(response.Status, response)
}
