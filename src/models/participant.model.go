package models

type Participant struct {
	UserId     uint `json:"user_id" gorm:"not null"`
	ChatRoomId uint `json:"chat_room_id" gorm:"not null"`
}
