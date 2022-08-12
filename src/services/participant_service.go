package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
)

type IParticipantService interface {
	CreateParticipant(int, uint) error
}

type ParticipantService struct {
	participantRepository repositories.IParticipantRepository
}

func NewParticipantService(participantRepository repositories.IParticipantRepository) *ParticipantService {
	return &ParticipantService{participantRepository}
}

func (participantService *ParticipantService) CreateParticipant(userId int, chatRoomId uint) error {
	participant := &models.Participant{
		UserId:     uint(userId),
		ChatRoomId: chatRoomId,
	}

	err := participantService.participantRepository.CreateParticipant(participant)

	return err
}
