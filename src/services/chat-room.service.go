package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/schemas"
)

type IChatRoomService interface {
	FindChatRoomById(int) (models.ChatRoom, error)
	CreateChatRoom(models.ChatRoom) (models.ChatRoom, error)
	UpdateChatRoom(models.ChatRoom, int) error
	RemoveChatRoom(int) error

	FindAllChatRoomByUserId(int) ([]models.ChatRoom, error)
	FindAllParticipantByChatRoomId(int) ([]models.Participant, error)
}

type ChatRoomService struct {
	chatRoomRepository    *repositories.ChatRoomRepository
	participantRepository *repositories.ParticipantRepository
}

func NewChatRoomService(chatRoomRepository *repositories.ChatRoomRepository, participantRepository *repositories.ParticipantRepository) *ChatRoomService {
	return &ChatRoomService{chatRoomRepository, participantRepository}
}

func (chatRoomService *ChatRoomService) FindChatRoomById(chatRoomId int) (models.ChatRoom, error) {
	chatRoom, err := chatRoomService.chatRoomRepository.FindChatRoomById(chatRoomId)

	return chatRoom, err
}

func (chatRoomService *ChatRoomService) CreateChatRoom(chatRoomRequest schemas.ChatRoomRequest) (models.ChatRoom, error) {
	request := models.ChatRoom{
		ImageUrl:    chatRoomRequest.ImageUrl,
		Description: chatRoomRequest.Description,
		Name:        chatRoomRequest.Name,
		Type:        chatRoomRequest.Type,
	}

	chatRoom, err := chatRoomService.chatRoomRepository.CreateChatRoom(request)

	return chatRoom, err
}

func (chatRoomService *ChatRoomService) UpdateChatRoom(chatRoomRequest *schemas.ChatRoomRequest, chatRoomId int) error {
	chatRoom := &models.ChatRoom{
		Description: chatRoomRequest.Description,
		ImageUrl:    chatRoomRequest.ImageUrl,
		Name:        chatRoomRequest.Name,
		Type:        chatRoomRequest.Type,
	}

	err := chatRoomService.chatRoomRepository.UpdateChatRoomById(chatRoom, chatRoomId)

	return err
}

func (chatRoomService *ChatRoomService) RemoveChatRoom(chatRoomId int) error {
	err := chatRoomService.chatRoomRepository.RemoveChatRoomById(chatRoomId)

	return err
}

func (chatRoomService *ChatRoomService) FindAllParticipantByChatRoomId(chatRoomId int) ([]uint, error) {
	participants, err := chatRoomService.participantRepository.FindAllParticipant(chatRoomId)

	var data []uint

	for _, participant := range participants {
		data = append(data, participant.UserId)
	}

	return data, err
}

func (chatRoomService *ChatRoomService) FindAllChatRoomByUserId(userId int) ([]schemas.ChatRoomWithPartisipants, error) {
	participants, err := chatRoomService.participantRepository.FindAllChatRoom(userId)

	var chatRooms []schemas.ChatRoomWithPartisipants

	for _, participant := range participants {
		chatRoom, _ := chatRoomService.FindChatRoomById(int(participant.ChatRoomId))
		chatRoomMembers, _ := chatRoomService.FindAllParticipantByChatRoomId(int(participant.ChatRoomId))

		chatRooms = append(chatRooms, schemas.ChatRoomWithPartisipants{
			ChatRoomId:  chatRoom.ChatRoomId,
			UserIds:     chatRoomMembers,
			ImageUrl:    chatRoom.ImageUrl,
			Description: chatRoom.Description,
			Name:        chatRoom.Name,
			Type:        chatRoom.Type,
			CreatedAt:   chatRoom.CreatedAt,
			UpdatedAt:   chatRoom.UpdatedAt,
		})
	}

	return chatRooms, err
}
