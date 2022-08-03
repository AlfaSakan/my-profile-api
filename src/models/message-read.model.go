package models

type MessageRead struct {
	MessageReadId uint  `json:"chat_room_message_id"  gorm:"primaryKey"`
	UserId        uint  `json:"sender_id"`
	ChatRoomId    uint  `json:"chat_room_id"`
	MessageId     uint  `json:"message_id"`
	IsRead        bool  `json:"is_read" gorm:"default:false"`
	CreatedAt     int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
