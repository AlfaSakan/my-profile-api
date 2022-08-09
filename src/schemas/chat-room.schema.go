package schemas

type ChatRoomRequest struct {
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	UserId      int    `json:"user_id"`
}

type ChatRoomWithPartisipants struct {
	ChatRoomId  uint   `json:"chat_room_id"`
	UserIds     []uint `json:"user_ids"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
