package utils

import (
	"time"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"

	"gorm.io/gorm"
)

func ReIssueAccessToken(db *gorm.DB, refreshToken string) (string, *models.User) {
	sessionRepository := repositories.NewSessionRepository(db)
	userRepository := repositories.NewUserRepository(db)

	claims, err := DecodeToken(refreshToken)
	if err != nil {
		return "", nil
	}

	data := claims["data"].(map[string]interface{})
	sessionId := data["session_id"].(string)

	session := &models.Session{}
	err = sessionRepository.FindSession(session, sessionId)
	if err != nil {
		return "", nil
	}

	user, err := userRepository.FindUserById(session.UserId)
	if err != nil {
		return "", nil
	}

	newClaims := &CustomClaim{
		User: &user,
	}

	expireAccessToken := time.Now().Add(time.Hour * 12).UnixMilli()
	accessToken, err := GenerateToken(newClaims, expireAccessToken)
	if err != nil {
		return "", nil
	}

	return accessToken, &user
}
