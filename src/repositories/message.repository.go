package repositories

import (
	"myProfileApi/src/models"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	FindMessagesByChatRoomId(int) ([]models.Message, error)
	FindMessageById(messageId int) (models.Message, error)
	CreateMessage(models.Message) (models.Message, error)
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

func (repository *MessageRepository) FindMessagesByChatRoomId(chatRoomId int) ([]models.Message, error) {
	var messages []models.Message

	err := repository.db.Where("ChatRoomId = ?", chatRoomId).Find(&messages).Error

	return messages, err
}
