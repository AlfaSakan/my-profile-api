package models

type Participant struct {
	UserId     string `json:"user_id" gorm:"not null"`
	ChatRoomId string `json:"chat_room_id" gorm:"not null"`
	UserStatus string `json:"user_status" gorm:"default:user"`
}
