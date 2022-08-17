package models

type Participant struct {
	UserId     uint   `json:"user_id" gorm:"not null"`
	ChatRoomId uint   `json:"chat_room_id" gorm:"not null"`
	UserStatus string `json:"user_status" gorm:"default:user"`
}
