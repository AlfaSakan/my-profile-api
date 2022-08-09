package schemas

type MessageRequest struct {
	ChatRoomId uint   `json:"chat_room_id" binding:"required"`
	UserId     uint   `json:"user_id" binding:"required"`
	Type       string `json:"type"`
	Message    string `json:"message" binding:"required"`
}

type MessageUpdate struct {
	StatusMessage string `json:"status" binding:"required"`
}
