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

	user, _ := s.userService.FindUser(&schemas.UserRequest{Name: request.Name, PhoneNumber: request.PhoneNumber})
	if len(user.Name) == 0 {
		utils.ResponseBadRequest(ctx, response, fmt.Errorf("user not found"))
		return
	}

	userAgent := ctx.Request.UserAgent()

	session, err := s.sessionService.Login(user.UserId, userAgent)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	accessClaims := &utils.CustomClaim{
		User: user,
	}

	expireAccessToken := time.Now().Add(time.Hour * 12).UnixMilli()
	accessToken, err := utils.GenerateToken(accessClaims, expireAccessToken)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	refreshClaim := &utils.CustomClaim{
		User:      &models.User{},
		SessionId: session.SessionId,
	}

	expireRefreshToken := time.Now().Add(time.Hour * 24 * 30 * 12).UnixMilli()
	refreshToken, err := utils.GenerateToken(refreshClaim, expireRefreshToken)
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

	sessionId := ctx.Param("sessionId")

	err := s.sessionService.Logout(sessionId)
	if err != nil {
		utils.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "Success Logout"
	response.Status = http.StatusOK

	ctx.JSON(response.Status, response)
}
