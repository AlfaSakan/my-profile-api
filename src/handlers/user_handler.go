package handlers

import (
	"net/http"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/services"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) GetUserHandler(ctx *gin.Context) {
	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	response := new(schemas.Response)

	user, errService := userHandler.userService.FindUserById(userId)
	if errService != nil {
		response.Message = errService.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = user
	ctx.JSON(http.StatusOK, response)
}

func (userHandler *UserHandler) PostUserHandler(ctx *gin.Context) {
	var request schemas.UserRequest
	response := new(schemas.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			response.Status = http.StatusBadRequest
			response.Message = e.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
	}

	responseService, errorService := userHandler.userService.CreateUser(request)
	if errorService != nil {
		response.Message = errorService.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "OK"
	response.Data = responseService
	ctx.JSON(http.StatusCreated, response)
}

func (userHandler *UserHandler) PatchUserHandler(ctx *gin.Context) {
	var request schemas.UpdateUserRequest
	response := new(schemas.Response)

	user, _ := ctx.Get("User")
	userId := user.(*models.User).UserId

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			response.Status = http.StatusBadRequest
			response.Message = e.Error()
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
	}

	_, errorService := userHandler.userService.UpdateUser(request, userId)
	if errorService != nil {
		response.Message = errorService.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = "Success Updated"
	ctx.JSON(http.StatusOK, response)
}

func (uh *UserHandler) SearchUserHandler(c *gin.Context) {
	name := c.Param("Name")
	response := &schemas.Response{}

	users, err := uh.userService.SearchUserByName(name)
	if err != nil {
		utils.ResponseBadRequest(c, response, err)
		return
	}

	response.Data = users
	response.Message = "OK"
	response.Status = http.StatusOK
	c.JSON(response.Status, response)
}

func (uh *UserHandler) GetOneUserHandler(c *gin.Context) {
	userId := c.Param("UserId")
	response := &schemas.Response{}

	user, err := uh.userService.FindUserById(userId)
	if err != nil {
		utils.ResponseBadRequest(c, response, err)
		return
	}

	response.Data = user
	response.Message = "OK"
	response.Status = http.StatusOK
	c.JSON(response.Status, response)
}
