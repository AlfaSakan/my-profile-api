package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/schemas"
)

type IChatRoomService interface {
	FindChatRoomById(int) (*schemas.ChatRoomWithPartisipants, error)
	CreateChatRoom(schemas.ChatRoomRequest) (*schemas.ChatRoomWithPartisipants, error)
	UpdateChatRoom(chatRoomRequest *schemas.ChatRoomRequest, chatRoomId int) error
	RemoveChatRoom(userId uint, chatRoomId int) error
	FindAllChatRoomByUserId(uint) ([]schemas.ChatRoomWithPartisipants, error)
	FindAllParticipantByChatRoomId(int) ([]uint, error)
}

type ChatRoomService struct {
	chatRoomRepository    repositories.IChatRoomRepository
	participantRepository repositories.IParticipantRepository
	messageRepository     repositories.IMessageRepository
}

func NewChatRoomService(cR repositories.IChatRoomRepository, pR repositories.IParticipantRepository, mR repositories.IMessageRepository) *ChatRoomService {
	return &ChatRoomService{chatRoomRepository: cR, participantRepository: pR, messageRepository: mR}
}

func (chatRoomService *ChatRoomService) FindChatRoomById(chatRoomId int) (*schemas.ChatRoomWithPartisipants, error) {
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

	var participantsId = []uint{}

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
	}
	var memberId = []uint{}

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
		UserId:     uint(chatRoomRequest.UserId),
		UserStatus: "admin",
		ChatRoomId: chatRoom.ChatRoomId,
	})

	if err != nil {
		return &schemas.ChatRoomWithPartisipants{}, err
	}

	memberId = append(memberId, uint(chatRoomRequest.UserId))

	for _, part := range chatRoomRequest.ParticipantsId {
		err = chatRoomService.participantRepository.CreateParticipant(&models.Participant{
			UserId:     uint(part),
			ChatRoomId: chatRoom.ChatRoomId,
		})

		if err != nil {
			continue
		}

		memberId = append(memberId, uint(part))
	}

	chatRoomCreated.ParticipantsId = &memberId

	return chatRoomCreated, err
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

func (chatRoomService *ChatRoomService) RemoveChatRoom(userId uint, chatRoomId int) error {
	err := chatRoomService.participantRepository.RemoveParticipant(userId, uint(chatRoomId))

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

func (cS *ChatRoomService) FindAllChatRoomByUserId(userId uint) ([]schemas.ChatRoomWithPartisipants, error) {
	participants, err := cS.participantRepository.FindAllChatRoom(userId)

	var chatRooms []schemas.ChatRoomWithPartisipants

	for _, participant := range participants {
		chatRoom, _ := cS.chatRoomRepository.FindChatRoomById(int(participant.ChatRoomId))
		chatRoomMembers, _ := cS.FindAllParticipantByChatRoomId(int(participant.ChatRoomId))

		messages, err := cS.messageRepository.FindMessagesByChatRoomId(int(chatRoom.ChatRoomId))

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
