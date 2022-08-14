package handlers

import (
	"fmt"
	"myProfileApi/src/models"
	"myProfileApi/src/schemas"
	"myProfileApi/src/services"
	"myProfileApi/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	sessionService services.ISessionService
	userService    services.IUserService
}

func NewSessionHandler(sessionService services.ISessionService, userService services.IUserService) *SessionHandler {
	return &SessionHandler{sessionService, userService}
}

func (s *SessionHandler) PostSessionHandler(ctx *gin.Context) {
	request := &schemas.SessionRequest{}

	response := &schemas.Response{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, fmt.Errorf("user_id cannot be empty"))
		return
	}

	user, err := s.userService.FindUserById(uint(request.UserId))
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	request.UserAgent = ctx.Request.UserAgent()

	session, err := s.sessionService.Login(request)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	accessClaims := &utils.CustomClaim{
		User: &user,
	}

	accessToken, err := utils.GenerateToken(accessClaims)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	refreshClaim := &utils.CustomClaim{
		User:      &models.User{},
		SessionId: session.SessionId,
	}

	refreshToken, err := utils.GenerateToken(refreshClaim)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = &models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		SessionId:    session.SessionId,
	}

	ctx.JSON(response.Status, response)
}

func (s *SessionHandler) DeleteSessionHandler(ctx *gin.Context) {
	response := &schemas.Response{}
	sessionId := utils.ConvertParamToInt(ctx, "sessionId")

	if sessionId == 0 {
		return
	}

	err := s.sessionService.Logout(sessionId)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "Success Logout"
	response.Status = http.StatusOK

	ctx.JSON(response.Status, response)
}
