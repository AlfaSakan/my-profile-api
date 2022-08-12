package handlers

import (
	"myProfileApi/src/schemas"
	"myProfileApi/src/services"
	"myProfileApi/src/utils"
	"net/http"

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
	userId := utils.ConvertParamToInt(ctx, "userId")

	if userId == 0 {
		return
	}

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

func (chatRoomHandler *ChatRoomHandler) PostChatRoom(ctx *gin.Context) {
	var request schemas.ChatRoomRequest
	response := new(schemas.Response)

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

	errParticipant := chatRoomHandler.participantService.CreateParticipant(request.UserId, chatRoom.ChatRoomId)
	if errParticipant != nil {
		response.Message = errParticipant.Error()
		response.Status = http.StatusBadRequest

		ctx.JSON(response.Status, response)
		return
	}

	response.Data = &chatRoom
	response.Message = "Created"
	response.Status = http.StatusCreated
	ctx.JSON(response.Status, response)
}

func (chatRoomHandler *ChatRoomHandler) PatchChatRoom(ctx *gin.Context) {
	var request schemas.ChatRoomRequest
	response := new(schemas.Response)

	chatRoomId := utils.ConvertParamToInt(ctx, "chatRoomId")

	if chatRoomId == 0 {
		return
	}

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

	chatRoomId := utils.ConvertParamToInt(ctx, "chatRoomId")

	if chatRoomId == 0 {
		return
	}

	err := chatRoomHandler.chatRoomService.RemoveChatRoom(chatRoomId)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusBadRequest

		ctx.JSON(response.Status, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = ""
	response.Message = "Success deleted"

	ctx.JSON(response.Status, response)
}
