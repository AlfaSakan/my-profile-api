package services

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"
)

type ISessionService interface {
	Login(userId string, userAgent string) (*models.Session, error)
	Logout(sessionId string) error
}

type SessionService struct {
	sessionRepository repositories.ISessionRepository
	userRepository    repositories.IUserRepository
}

func NewSessionService(sessionRepository repositories.ISessionRepository, userRepository repositories.IUserRepository) *SessionService {
	return &SessionService{sessionRepository, userRepository}
}

func (s *SessionService) Login(userId string, userAgent string) (*models.Session, error) {
	session := &models.Session{
		UserId:    userId,
		UserAgent: userAgent,
		SessionId: utils.GenerateId(),
	}

	err := s.sessionRepository.CreateSession(session)

	return session, err
}

func (s *SessionService) Logout(sessionId string) error {
	session := &models.Session{
		SessionId: sessionId,
	}

	return s.sessionRepository.DeleteSession(session, sessionId)
}
