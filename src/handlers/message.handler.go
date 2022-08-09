package handlers

import (
	"myProfileApi/src/schemas"
	"myProfileApi/src/services"
	"myProfileApi/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MessageHandler struct {
	messageService *services.MessageService
}

func NewMessageHandler(messageService *services.MessageService) *MessageHandler {
	return &MessageHandler{messageService}
}

func (messageHandler *MessageHandler) GetMessageHandler(ctx *gin.Context) {
	chatRoomId := ctx.Param("chatRoomId")
	response := new(schemas.Response)

	chatRoomIdInt, errConvert := strconv.Atoi(chatRoomId)
	if errConvert != nil {
		response.Message = errConvert.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	messages, err := messageHandler.messageService.FindMessageByChatRoomId(chatRoomIdInt)
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

	errRequest := ctx.ShouldBindJSON(&request)
	if errRequest != nil {
		for _, e := range errRequest.(validator.ValidationErrors) {
			response.Status = http.StatusBadRequest
			response.Message = e.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
	}

	responseService, err := messageHandler.messageService.CreateMessage(request)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status = http.StatusCreated
	response.Data = responseService
	response.Message = "OK"
	ctx.JSON(http.StatusCreated, response)
}

func (messageHandler *MessageHandler) PatchMessageHandler(ctx *gin.Context) {
	messageId := utils.ConvertParamToInt(ctx, "messageId")
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
