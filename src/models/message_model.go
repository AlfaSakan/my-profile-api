package models

type Message struct {
	MessageId     uint   `json:"message_id" gorm:"primaryKey"`
	ChatRoomId    uint   `json:"chat_room_id" gorm:"not null"`
	UserId        uint   `json:"user_id" gorm:"not null"`
	StatusMessage string `json:"status" gorm:"default:active;size:50"`
	Type          string `json:"type" gorm:"size:50;default:chat"`
	Message       string `json:"message"`
	CreatedAt     int64  `json:"created_at" gorm:"autoCreateTime:milli"`
}
