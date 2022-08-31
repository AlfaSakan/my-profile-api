package services

import (
	"fmt"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"
)

type IParticipantService interface {
	CreateParticipant(userId string, chatRoomId string) error
	RemoveParticipant(userId string, chatRoomId string) error
	FindUserAdmin(userId string, chatRoomId string) (*models.Participant, error)
	FindParticipantsList(chatRoomId string) (*[]models.User, error)
}

type ParticipantService struct {
	participantRepository repositories.IParticipantRepository
	userRepository        repositories.IUserRepository
}

func NewParticipantService(participantRepository repositories.IParticipantRepository, userRepository repositories.IUserRepository) *ParticipantService {
	return &ParticipantService{participantRepository, userRepository}
}

func (participantService *ParticipantService) CreateParticipant(userId string, chatRoomId string) error {
	participant := &models.Participant{
		UserId:     userId,
		ChatRoomId: chatRoomId,
	}

	err := participantService.participantRepository.CreateParticipant(participant)

	return err
}

func (participantService *ParticipantService) RemoveParticipant(userId string, chatRoomId string) error {
	return participantService.participantRepository.RemoveParticipant(userId, chatRoomId)
}

func (participantService *ParticipantService) FindUserAdmin(userId string, chatRoomId string) (*models.Participant, error) {
	participant, err := participantService.participantRepository.FindOne(userId, chatRoomId)

	if participant.UserStatus != "admin" {
		return nil, fmt.Errorf("not admin")
	}

	return participant, err
}

func (participantService *ParticipantService) FindParticipantsList(chatRoomId string) (*[]models.User, error) {
	participants, err := participantService.participantRepository.FindAllParticipant(chatRoomId)

	users := []models.User{}

	for _, participant := range participants {
		userFound, _ := participantService.userRepository.FindUserById(participant.UserId)

		users = append(users, userFound)
	}

	return &users, err
}
