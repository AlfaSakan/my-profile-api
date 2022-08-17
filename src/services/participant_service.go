package services

import (
	"fmt"
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
)

type IParticipantService interface {
	CreateParticipant(userId int, chatRoomId uint) error
	RemoveParticipant(userId int, chatRoomId int) error
	FindUserAdmin(userId uint, chatRoomId int) (*models.Participant, error)
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

func (participantService *ParticipantService) RemoveParticipant(userId int, chatRoomId int) error {
	return participantService.participantRepository.RemoveParticipant(uint(userId), uint(chatRoomId))
}

func (participantService *ParticipantService) FindUserAdmin(userId uint, chatRoomId int) (*models.Participant, error) {
	participant, err := participantService.participantRepository.FindOne(uint(userId), uint(chatRoomId))

	if participant.UserStatus != "admin" {
		return nil, fmt.Errorf("not admin")
	}

	return participant, err
}
