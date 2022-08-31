package schemas

type MessageRequest struct {
	ChatRoomId string `json:"chat_room_id" binding:"required"`
	SenderId   string `json:"sender_id" binding:"required"`
	Type       string `json:"type"`
	Message    string `json:"message" binding:"required"`
}

type MessageUpdate struct {
	StatusMessage string `json:"status" binding:"required"`
}
