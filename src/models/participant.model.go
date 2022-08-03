package models

type Participant struct {
	ParticipantId uint `json:"participant_id" gorm:"primaryKey"`
	UserId        uint `json:"user_id" gorm:"not null"`
	ChatRoomId    uint `json:"chat_room_id" gorm:"not null"`
}
