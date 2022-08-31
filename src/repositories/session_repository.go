package repositories

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"gorm.io/gorm"
)

type ISessionRepository interface {
	CreateSession(session *models.Session) error
	FindSession(session *models.Session, sessionId string) error
	DeleteSession(session *models.Session, sessionId string) error
}

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) FindSession(session *models.Session, sessionId string) error {
	return r.db.Where(&models.Session{SessionId: sessionId}).Find(session).Error
}

func (r *SessionRepository) CreateSession(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepository) DeleteSession(session *models.Session, sessionId string) error {
	return r.db.Delete(session, sessionId).Error
}
