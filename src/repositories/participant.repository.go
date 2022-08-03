package repositories

import (
	"myProfileApi/src/models"

	"gorm.io/gorm"
)

type IParticipantRepository interface {
	GetAllParticipant(int) ([]models.Participant, error)
	GetAllChatRoom(int) ([]models.Participant, error)
}

type ParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) *ParticipantRepository {
	return &ParticipantRepository{db}
}

func (participantRepository *ParticipantRepository) GetAllParticipant(chatRoomId int) ([]models.Participant, error) {
	var participants []models.Participant

	error := participantRepository.db.Where(&models.Participant{ChatRoomId: uint(chatRoomId)}).Find(&participants).Error

	return participants, error
}

func (participantRepository *ParticipantRepository) GetAllChatRoom(userId int) ([]models.Participant, error) {
	var participants []models.Participant

	error := participantRepository.db.Where(&models.Participant{ChatRoomId: uint(userId)}).Find(&participants).Error

	return participants, error
}
