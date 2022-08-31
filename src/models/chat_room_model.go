package models

type ChatRoom struct {
	ChatRoomId  string `json:"chat_room_id" gorm:"primaryKey"`
	ImageUrl    string `json:"image_url" gorm:"size:200"`
	Description string `json:"description" gorm:"size:200"`
	Name        string `json:"name" gorm:"size:50"`
	Type        string `json:"type"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
