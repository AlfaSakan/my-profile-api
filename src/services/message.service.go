package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/schemas"
)

type IMessageService interface {
	CreateMessage(schemas.MessageRequest) (models.Message, error)
	FindMessageByChatRoomId(int) ([]models.Message, error)
	UpdateMessage(*models.Message, int) error
}

type MessageService struct {
	messageRepository *repositories.MessageRepository
}

func NewMessageService(messageRepository *repositories.MessageRepository) *MessageService {
	return &MessageService{messageRepository}
}

func (messageService *MessageService) FindMessageByChatRoomId(chatRoomId int) (*[]models.Message, error) {
	messages, err := messageService.messageRepository.FindMessagesByChatRoomId(chatRoomId)

	return messages, err
}

func (messageService *MessageService) CreateMessage(messageRequest schemas.MessageRequest) (models.Message, error) {
	data := models.Message{
		Message:    messageRequest.Message,
		UserId:     messageRequest.UserId,
		ChatRoomId: messageRequest.ChatRoomId,
		Type:       messageRequest.Type,
	}

	message, err := messageService.messageRepository.CreateMessage(data)

	return message, err
}

func (messageService *MessageService) UpdateMessage(messageRequest *schemas.MessageUpdate, messageId int) error {
	message := &models.Message{
		StatusMessage: messageRequest.StatusMessage,
	}

	err := messageService.messageRepository.UpdateMessagesByMessageId(message, messageId)

	return err
}
