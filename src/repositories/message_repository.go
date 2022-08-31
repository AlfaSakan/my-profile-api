package repositories

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	FindMessagesByChatRoomId(chatRoomId string) (*[]models.Message, error)
	FindMessageById(messageId int) (models.Message, error)
	CreateMessage(models.Message) (models.Message, error)
	UpdateMessagesByMessageId(messageRequest *models.Message, messageId string) error
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

func (repository *MessageRepository) FindMessagesByChatRoomId(chatRoomId string) (*[]models.Message, error) {
	var messages []models.Message
	chatRoom := &models.ChatRoom{
		ChatRoomId: chatRoomId,
	}

	err := repository.db.Order("created_at DESC").Where(chatRoom).Find(&messages).Error

	return &messages, err
}

func (repository *MessageRepository) UpdateMessagesByMessageId(messageRequest *models.Message, messageId string) error {
	message := &models.Message{
		MessageId: messageId,
	}

	err := repository.db.Where(&message).Updates(messageRequest).Error

	return err
}
