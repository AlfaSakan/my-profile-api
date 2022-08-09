package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
)

type IParticipantService interface {
	CreateParticipant(*models.Participant) error
}

type ParticipantService struct {
	participantRepository *repositories.ParticipantRepository
}

func NewParticipantService(participantRepository *repositories.ParticipantRepository) *ParticipantService {
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
