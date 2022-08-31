package repositories

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"gorm.io/gorm"
)

type IChatRoomRepository interface {
	CreateChatRoom(*models.ChatRoom) (*models.ChatRoom, error)
	FindChatRoomById(chatRoomId string) (models.ChatRoom, error)
	RemoveChatRoomById(chatRoomId string) error
	UpdateChatRoomById(chatRoomRequest *models.ChatRoom, chatRoomId string) error
}

type ChatRoomRepository struct {
	db *gorm.DB
}

func NewChatRoomRepository(db *gorm.DB) *ChatRoomRepository {
	return &ChatRoomRepository{db}
}

func (chatRoomRepository *ChatRoomRepository) CreateChatRoom(chatRoom *models.ChatRoom) (*models.ChatRoom, error) {
	err := chatRoomRepository.db.Create(&chatRoom).Error

	return chatRoom, err
}

func (chatRoomRepository *ChatRoomRepository) FindChatRoomById(chatRoomId string) (models.ChatRoom, error) {
	var chatRoom models.ChatRoom

	err := chatRoomRepository.db.Where(&models.ChatRoom{ChatRoomId: chatRoomId}).First(&chatRoom).Error

	return chatRoom, err
}

func (chatRoomRepository *ChatRoomRepository) RemoveChatRoomById(chatRoomId string) error {
	chatRoom := &models.ChatRoom{
		ChatRoomId: chatRoomId,
	}

	err := chatRoomRepository.db.Delete(chatRoom).Error

	return err
}

func (chatRoomRepository *ChatRoomRepository) UpdateChatRoomById(chatRoomRequest *models.ChatRoom, chatRoomId string) error {
	chatRoom := &models.ChatRoom{
		ChatRoomId: chatRoomId,
	}

	err := chatRoomRepository.db.Where(chatRoom).Updates(chatRoomRequest).Error

	return err
}
