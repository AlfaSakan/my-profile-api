package repositories

import (
	"myProfileApi/src/models"

	"gorm.io/gorm"
)

type IParticipantRepository interface {
	FindAllParticipant(int) ([]models.Participant, error)
	FindAllChatRoom(uint) ([]models.Participant, error)
	CreateParticipant(*models.Participant) error
}

type ParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) *ParticipantRepository {
	return &ParticipantRepository{db}
}

func (participantRepository *ParticipantRepository) FindAllParticipant(chatRoomId int) ([]models.Participant, error) {
	var participants []models.Participant

	err := participantRepository.db.Where(&models.Participant{ChatRoomId: uint(chatRoomId)}).Find(&participants).Error

	return participants, err
}

func (participantRepository *ParticipantRepository) FindAllChatRoom(userId uint) ([]models.Participant, error) {
	var participants []models.Participant

	err := participantRepository.db.Where(&models.Participant{UserId: userId}).Find(&participants).Error

	return participants, err
}

func (participantRepository *ParticipantRepository) CreateParticipant(participant *models.Participant) error {
	err := participantRepository.db.Create(participant).Error

	return err
}
