package handlers

import (
	"fmt"
	"myProfileApi/src/models"
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

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId
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

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	request.UserId = int(userId)

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

	response.Data = &chatRoom
	response.Message = "Created"
	response.Status = http.StatusCreated
	ctx.JSON(response.Status, response)
}

func (chatRoomHandler *ChatRoomHandler) PatchChatRoom(ctx *gin.Context) {
	var request schemas.ChatRoomRequest
	response := new(schemas.Response)

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	request.UserId = int(userId)

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
	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	chatRoomId := utils.ConvertParamToInt(ctx, "chatRoomId")

	if chatRoomId == 0 {
		utils.ResponseBadRequest(ctx, response, fmt.Errorf("chat room id error"))
		return
	}

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

func (ChatRoomHandler *ChatRoomHandler) PostAddParticipant(ctx *gin.Context) {
	request := &schemas.AddParticipantRequest{}
	response := &schemas.Response{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	participant, err := ChatRoomHandler.participantService.FindUserAdmin(userId, request.ChatRoomId)

	if participant == nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	for _, userId := range request.UserIds {
		ChatRoomHandler.participantService.CreateParticipant(userId, uint(request.ChatRoomId))
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = ""
	ctx.JSON(response.Status, response)
}

func (ChatRoomHandler *ChatRoomHandler) GetChatRoomById(ctx *gin.Context) {
	response := &schemas.Response{}
	chatRoomId := utils.ConvertParamToInt(ctx, "chatRoomId")

	if chatRoomId == 0 {
		utils.ResponseBadRequest(ctx, response, fmt.Errorf("error chat room id"))
		return
	}

	chatRoom, err := ChatRoomHandler.chatRoomService.FindChatRoomById(chatRoomId)

	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = chatRoom
	ctx.JSON(response.Status, response)
}
