package schemas

type MessageRequest struct {
	ChatRoomId uint   `json:"chat_room_id"`
	UserId     uint   `json:"user_id" gorm:"not null"`
	Type       string `json:"type"`
	Message    string `json:"message"`
}
