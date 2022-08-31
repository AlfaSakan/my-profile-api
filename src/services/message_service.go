package services

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"
)

type IMessageService interface {
	CreateMessage(schemas.MessageRequest) (*models.Message, error)
	FindMessageByChatRoomId(chatRoomId string) (*[]models.Message, error)
	UpdateMessage(messageRequest *schemas.MessageUpdate, messageId string) error
}

type MessageService struct {
	messageRepository repositories.IMessageRepository
}

func NewMessageService(messageRepository repositories.IMessageRepository) *MessageService {
	return &MessageService{messageRepository}
}

func (messageService *MessageService) FindMessageByChatRoomId(chatRoomId string) (*[]models.Message, error) {
	messages, err := messageService.messageRepository.FindMessagesByChatRoomId(chatRoomId)

	return messages, err
}

func (messageService *MessageService) CreateMessage(messageRequest schemas.MessageRequest) (*models.Message, error) {
	data := models.Message{
		Message:    messageRequest.Message,
		SenderId:   messageRequest.SenderId,
		ChatRoomId: messageRequest.ChatRoomId,
		Type:       messageRequest.Type,
		MessageId:  utils.GenerateId(),
	}

	message, err := messageService.messageRepository.CreateMessage(data)

	return &message, err
}

func (messageService *MessageService) UpdateMessage(messageRequest *schemas.MessageUpdate, messageId string) error {
	message := &models.Message{
		StatusMessage: messageRequest.StatusMessage,
	}

	err := messageService.messageRepository.UpdateMessagesByMessageId(message, messageId)

	return err
}
