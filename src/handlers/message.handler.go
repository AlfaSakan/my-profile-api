package handlers

import (
	"myProfileApi/src/schemas"
	"myProfileApi/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IMessageHandler interface {
}

type MessageHandler struct {
	messageService *services.MessageService
}

func NewMessageHandler(messageService *services.MessageService) *MessageHandler {
	return &MessageHandler{messageService}
}

func (messageHandler *MessageHandler) GetMessageHandler(ctx *gin.Context) {
	chatRoomId := ctx.Param("chatRoomId")

	chatRoomIdInt, errConvert := strconv.Atoi(chatRoomId)
	if errConvert != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response{
			ErrorMessage: errConvert.Error(),
			Status:       http.StatusBadRequest,
			Data:         "",
		})
		return
	}

	messages, err := messageHandler.messageService.FindMessageByChatRoomId(chatRoomIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response{
			ErrorMessage: err.Error(),
			Status:       http.StatusBadRequest,
			Data:         "",
		})
		return
	}

	ctx.JSON(http.StatusOK, schemas.Response{
		ErrorMessage: "",
		Status:       http.StatusOK,
		Data:         messages,
	})
}

func (messageHandler *MessageHandler) PostMessageHandler(ctx *gin.Context) {
	var request schemas.MessageRequest

	errRequest := ctx.ShouldBindJSON(&request)
	if errRequest != nil {
		for _, e := range errRequest.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, schemas.Response{
				Status:       http.StatusBadRequest,
				ErrorMessage: e.Error(),
				Data:         "",
			})
			return
		}
	}

	response, err := messageHandler.messageService.CreateMessage(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response{
			ErrorMessage: err.Error(),
			Status:       http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.Response{
		ErrorMessage: "",
		Status:       http.StatusCreated,
		Data:         response,
	})
}

func (messageHandler *MessageHandler) PatchMessageHandler(ctx *gin.Context) {}
