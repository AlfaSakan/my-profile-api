package repositories

import (
	"myProfileApi/src/models"

	"gorm.io/gorm"
)

type IChatRoomRepository interface {
	CreateChatRoom(*models.ChatRoom) (*models.ChatRoom, error)
	FindChatRoomById(int) (models.ChatRoom, error)
	RemoveChatRoomById(int) error
	UpdateChatRoomById(*models.ChatRoom, int) error
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

func (chatRoomRepository *ChatRoomRepository) FindChatRoomById(chatRoomId int) (models.ChatRoom, error) {
	var chatRoom models.ChatRoom

	err := chatRoomRepository.db.First(&chatRoom, chatRoomId).Error

	return chatRoom, err
}

func (chatRoomRepository *ChatRoomRepository) RemoveChatRoomById(chatRoomId int) error {
	chatRoom := &models.ChatRoom{
		ChatRoomId: uint(chatRoomId),
	}

	err := chatRoomRepository.db.Delete(chatRoom).Error

	return err
}

func (chatRoomRepository *ChatRoomRepository) UpdateChatRoomById(chatRoomRequest *models.ChatRoom, chatRoomId int) error {
	chatRoom := &models.ChatRoom{
		ChatRoomId: uint(chatRoomId),
	}

	err := chatRoomRepository.db.Where(chatRoom).Updates(chatRoomRequest).Error

	return err
}
