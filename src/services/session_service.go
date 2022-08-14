package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/schemas"
)

type ISessionService interface {
	Login(*schemas.SessionRequest) (*models.Session, error)
	Logout(int) error
}

type SessionService struct {
	sessionRepository repositories.ISessionRepository
	userRepository    repositories.IUserRepository
}

func NewSessionService(sessionRepository repositories.ISessionRepository, userRepository repositories.IUserRepository) *SessionService {
	return &SessionService{sessionRepository, userRepository}
}

func (s *SessionService) Login(sessionRequest *schemas.SessionRequest) (*models.Session, error) {
	session := &models.Session{
		UserId: uint(sessionRequest.UserId),
	}

	err := s.sessionRepository.CreateSession(session)

	return session, err
}

func (s *SessionService) Logout(sessionId int) error {
	session := &models.Session{
		SessionId: uint(sessionId),
	}

	return s.sessionRepository.DeleteSession(session, sessionId)
}
