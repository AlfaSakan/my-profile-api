package models

type MessageRead struct {
	UserId     uint `json:"sender_id"`
	ChatRoomId uint `json:"chat_room_id"`
	MessageId  uint `json:"message_id"`
	IsRead     bool `json:"is_read" gorm:"default:false"`
}
