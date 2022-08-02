package handlers

import (
	"myProfileApi/src/models"
	"myProfileApi/src/schemas"
	"myProfileApi/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) GetUserHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")
	userIdInt, errConvert := strconv.Atoi(userId)
	if errConvert != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response[models.User]{
			ErrorMessage: errConvert.Error(),
			Status:       http.StatusBadRequest,
		})
		return
	}

	user, errService := userHandler.userService.FindUserById(userIdInt)
	if errService != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response[models.User]{
			ErrorMessage: errService.Error(),
			Status:       http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, schemas.Response[models.User]{
		ErrorMessage: "",
		Status:       http.StatusOK,
		Data:         user,
	})
}

func (userHandler *UserHandler) PostUserHandler(ctx *gin.Context) {
	var request schemas.UserRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			// errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			// ctx.JSON(http.StatusBadRequest, errorMessage)
			ctx.JSON(http.StatusBadRequest, schemas.Response[string]{
				Status:       http.StatusBadRequest,
				ErrorMessage: e.Error(),
				Data:         "",
			})
			return
		}
	}

	response, errorService := userHandler.userService.CreateUser(request)
	if errorService != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response[models.User]{
			ErrorMessage: errorService.Error(),
			Status:       http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.Response[models.User]{
		Status:       http.StatusCreated,
		ErrorMessage: "",
		Data:         response,
	})
}

func (userHandler *UserHandler) PatchUserHandler(ctx *gin.Context) {
	var request schemas.UpdateUserRequest

	userId := ctx.Param("userId")
	userIdInt, errConvert := strconv.Atoi(userId)
	if errConvert != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response[models.User]{
			ErrorMessage: errConvert.Error(),
			Status:       http.StatusBadRequest,
		})
		return
	}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, schemas.Response[string]{
				Status:       http.StatusBadRequest,
				ErrorMessage: e.Error(),
				Data:         "",
			})
			return
		}
	}

	response, errorService := userHandler.userService.UpdateUser(request, userIdInt)
	if errorService != nil {
		ctx.JSON(http.StatusBadRequest, schemas.Response[models.User]{
			ErrorMessage: errorService.Error(),
			Status:       http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, schemas.Response[models.User]{
		Status:       http.StatusOK,
		ErrorMessage: "",
		Data:         response,
	})
}
