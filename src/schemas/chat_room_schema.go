package schemas

import "myProfileApi/src/models"

type ChatRoomRequest struct {
	ImageUrl       string `json:"image_url"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	UserId         int    `json:"user_id"`
	ParticipantsId []int  `json:"participants_id" binding:"required"`
	UpdatedAt      int64  `json:"updated_at"`
}

type ChatRoomWithPartisipants struct {
	ChatRoomId     uint              `json:"chat_room_id"`
	ParticipantsId *[]uint           `json:"participants_id"`
	Messages       *[]models.Message `json:"messages"`
	ImageUrl       string            `json:"image_url"`
	Description    string            `json:"description"`
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	CreatedAt      int64             `json:"created_at"`
	UpdatedAt      int64             `json:"updated_at"`
}

type AddParticipantRequest struct {
	ChatRoomId int   `json:"chat_room_id" binding:"required"`
	UserIds    []int `json:"user_ids" binding:"required"`
}
