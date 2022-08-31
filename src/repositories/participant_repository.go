package repositories

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"gorm.io/gorm"
)

type IParticipantRepository interface {
	FindAllParticipant(chatRoomId string) ([]models.Participant, error)
	FindAllChatRoom(userId string) ([]models.Participant, error)
	CreateParticipant(*models.Participant) error
	RemoveParticipant(userId string, chatRoomId string) error
	FindOne(userId string, chatRoomId string) (*models.Participant, error)
}

type ParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) *ParticipantRepository {
	return &ParticipantRepository{db}
}

func (participantRepository *ParticipantRepository) FindAllParticipant(chatRoomId string) ([]models.Participant, error) {
	var participants []models.Participant

	err := participantRepository.db.Where(&models.Participant{ChatRoomId: chatRoomId}).Find(&participants).Error

	return participants, err
}

func (participantRepository *ParticipantRepository) FindAllChatRoom(userId string) ([]models.Participant, error) {
	var participants []models.Participant

	err := participantRepository.db.Where(&models.Participant{UserId: userId}).Find(&participants).Error

	return participants, err
}

func (participantRepository *ParticipantRepository) CreateParticipant(participant *models.Participant) error {
	err := participantRepository.db.Create(participant).Error

	return err
}

func (participantRepository *ParticipantRepository) RemoveParticipant(userId string, chatRoomId string) error {
	return participantRepository.db.Where(&models.Participant{UserId: userId, ChatRoomId: chatRoomId}).Delete(&models.Participant{}).Error
}

func (participantRepository *ParticipantRepository) FindOne(userId string, chatRoomId string) (*models.Participant, error) {
	participant := &models.Participant{}

	err := participantRepository.db.Where(&models.Participant{UserId: userId, ChatRoomId: chatRoomId}).Find(participant).Error

	return participant, err
}
