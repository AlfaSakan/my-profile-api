package repositories

import (
	"myProfileApi/src/models"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	FindMessagesByChatRoomId(int) (*[]models.Message, error)
	FindMessageById(messageId int) (models.Message, error)
	CreateMessage(models.Message) (models.Message, error)
	UpdateMessagesByMessageId(*models.Message, int) error
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (repository *MessageRepository) FindMessageById(messageId int) (models.Message, error) {
	var message models.Message

	err := repository.db.Find(&message, messageId).Error

	return message, err
}

func (repository *MessageRepository) CreateMessage(message models.Message) (models.Message, error) {
	err := repository.db.Create(&message).Error

	return message, err
}

func (repository *MessageRepository) FindMessagesByChatRoomId(chatRoomId int) (*[]models.Message, error) {
	var messages []models.Message
	chatRoom := &models.ChatRoom{
		ChatRoomId: uint(chatRoomId),
	}

	err := repository.db.Where(chatRoom).Find(&messages).Error

	return &messages, err
}

func (repository *MessageRepository) UpdateMessagesByMessageId(messageRequest *models.Message, messageId int) error {
	message := &models.Message{
		MessageId: uint(messageId),
	}

	err := repository.db.Where(&message).Updates(messageRequest).Error

	return err
}
