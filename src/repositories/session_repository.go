package repositories

import (
	"myProfileApi/src/models"

	"gorm.io/gorm"
)

type ISessionRepository interface {
	CreateSession(*models.Session) error
	FindSession(*models.Session, int) error
	DeleteSession(*models.Session, int) error
}

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) FindSession(session *models.Session, sessionId int) error {
	return r.db.Find(session, sessionId).Error
}

func (r *SessionRepository) CreateSession(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepository) DeleteSession(session *models.Session, sessionId int) error {
	return r.db.Delete(session, sessionId).Error
}
