package utils

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"

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
	sessionId := int(data["session_id"].(float64))

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

	accessToken, err := GenerateToken(newClaims)
	if err != nil {
		return "", nil
	}

	return accessToken, &user
}
