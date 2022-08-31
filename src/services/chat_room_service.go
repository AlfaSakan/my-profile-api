package services

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"
)

type IChatRoomService interface {
	FindChatRoomById(chatRoomId string) (*schemas.ChatRoomWithPartisipants, error)
	CreateChatRoom(chatRoomRequest schemas.ChatRoomRequest) (*schemas.ChatRoomWithPartisipants, error)
	UpdateChatRoom(chatRoomRequest *schemas.ChatRoomRequest, chatRoomId string) error
	RemoveChatRoom(userId string, chatRoomId string) error
	FindAllChatRoomByUserId(userId string) ([]schemas.ChatRoomWithPartisipants, error)
	FindAllParticipantByChatRoomId(chatRoomId string) ([]string, error)
}

type ChatRoomService struct {
	chatRoomRepository    repositories.IChatRoomRepository
	participantRepository repositories.IParticipantRepository
	messageRepository     repositories.IMessageRepository
}

func NewChatRoomService(cR repositories.IChatRoomRepository, pR repositories.IParticipantRepository, mR repositories.IMessageRepository) *ChatRoomService {
	return &ChatRoomService{chatRoomRepository: cR, participantRepository: pR, messageRepository: mR}
}

func (chatRoomService *ChatRoomService) FindChatRoomById(chatRoomId string) (*schemas.ChatRoomWithPartisipants, error) {
	chatRoom, err := chatRoomService.chatRoomRepository.FindChatRoomById(chatRoomId)
	if err != nil {
		return &schemas.ChatRoomWithPartisipants{}, err
	}

	chatRoomFound := &schemas.ChatRoomWithPartisipants{
		ChatRoomId:  chatRoom.ChatRoomId,
		ImageUrl:    chatRoom.ImageUrl,
		Description: chatRoom.Description,
		CreatedAt:   chatRoom.CreatedAt,
		UpdatedAt:   chatRoom.UpdatedAt,
		Type:        chatRoom.Type,
		Name:        chatRoom.Name,
	}

	var participantsId = []string{}

	participants, err := chatRoomService.participantRepository.FindAllParticipant(chatRoomId)
	if err != nil {
		return chatRoomFound, err
	}
	for _, part := range participants {
		participantsId = append(participantsId, part.UserId)
	}

	messages, err := chatRoomService.messageRepository.FindMessagesByChatRoomId(chatRoomId)
	if err != nil {
		return chatRoomFound, err
	}

	chatRoomFound.ParticipantsId = &participantsId
	chatRoomFound.Messages = messages

	return chatRoomFound, err
}

func (chatRoomService *ChatRoomService) CreateChatRoom(chatRoomRequest schemas.ChatRoomRequest) (*schemas.ChatRoomWithPartisipants, error) {
	request := &models.ChatRoom{
		ImageUrl:    chatRoomRequest.ImageUrl,
		Description: chatRoomRequest.Description,
		Name:        chatRoomRequest.Name,
		Type:        chatRoomRequest.Type,
		ChatRoomId:  utils.GenerateId(),
	}
	var memberId = []string{}

	chatRoom, err := chatRoomService.chatRoomRepository.CreateChatRoom(request)

	if err != nil {
		return &schemas.ChatRoomWithPartisipants{}, err
	}

	chatRoomCreated := &schemas.ChatRoomWithPartisipants{
		ChatRoomId:  chatRoom.ChatRoomId,
		ImageUrl:    chatRoom.ImageUrl,
		Description: chatRoom.Description,
		Name:        chatRoom.Name,
		Type:        chatRoom.Type,
		CreatedAt:   chatRoom.CreatedAt,
		UpdatedAt:   chatRoom.UpdatedAt,
	}

	err = chatRoomService.participantRepository.CreateParticipant(&models.Participant{
		UserId:     chatRoomRequest.UserId,
		UserStatus: "admin",
		ChatRoomId: chatRoom.ChatRoomId,
	})

	if err != nil {
		return &schemas.ChatRoomWithPartisipants{}, err
	}

	memberId = append(memberId, chatRoomRequest.UserId)

	for _, part := range chatRoomRequest.ParticipantsId {
		err = chatRoomService.participantRepository.CreateParticipant(&models.Participant{
			UserId:     part,
			ChatRoomId: chatRoom.ChatRoomId,
		})

		if err != nil {
			continue
		}

		memberId = append(memberId, part)
	}

	chatRoomCreated.ParticipantsId = &memberId

	return chatRoomCreated, err
}

func (chatRoomService *ChatRoomService) UpdateChatRoom(chatRoomRequest *schemas.ChatRoomRequest, chatRoomId string) error {
	chatRoom := &models.ChatRoom{
		Description: chatRoomRequest.Description,
		ImageUrl:    chatRoomRequest.ImageUrl,
		Name:        chatRoomRequest.Name,
		Type:        chatRoomRequest.Type,
	}

	err := chatRoomService.chatRoomRepository.UpdateChatRoomById(chatRoom, chatRoomId)

	return err
}

func (chatRoomService *ChatRoomService) RemoveChatRoom(userId string, chatRoomId string) error {
	err := chatRoomService.participantRepository.RemoveParticipant(userId, chatRoomId)

	return err
}

func (chatRoomService *ChatRoomService) FindAllParticipantByChatRoomId(chatRoomId string) ([]string, error) {
	participants, err := chatRoomService.participantRepository.FindAllParticipant(chatRoomId)

	var data []string

	for _, participant := range participants {
		data = append(data, participant.UserId)
	}

	return data, err
}

func (cS *ChatRoomService) FindAllChatRoomByUserId(userId string) ([]schemas.ChatRoomWithPartisipants, error) {
	participants, err := cS.participantRepository.FindAllChatRoom(userId)

	var chatRooms []schemas.ChatRoomWithPartisipants

	for _, participant := range participants {
		chatRoom, _ := cS.chatRoomRepository.FindChatRoomById(participant.ChatRoomId)
		chatRoomMembers, _ := cS.FindAllParticipantByChatRoomId(participant.ChatRoomId)

		messages, err := cS.messageRepository.FindMessagesByChatRoomId(chatRoom.ChatRoomId)

		if err != nil {
			continue
		}

		chatRooms = append(chatRooms, schemas.ChatRoomWithPartisipants{
			ChatRoomId:     chatRoom.ChatRoomId,
			ParticipantsId: &chatRoomMembers,
			ImageUrl:       chatRoom.ImageUrl,
			Description:    chatRoom.Description,
			Name:           chatRoom.Name,
			Type:           chatRoom.Type,
			CreatedAt:      chatRoom.CreatedAt,
			UpdatedAt:      chatRoom.UpdatedAt,
			Messages:       messages,
		})
	}

	return chatRooms, err
}
