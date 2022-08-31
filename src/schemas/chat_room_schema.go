package schemas

import "github.com/AlfaSakan/my-profile-api.git/src/models"

type ChatRoomRequest struct {
	ImageUrl       string   `json:"image_url"`
	Description    string   `json:"description"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	UserId         string   `json:"user_id"`
	ParticipantsId []string `json:"participants_id" binding:"required"`
	UpdatedAt      int64    `json:"updated_at"`
}

type ChatRoomWithPartisipants struct {
	ChatRoomId     string            `json:"chat_room_id"`
	ParticipantsId *[]string         `json:"participants_id"`
	Messages       *[]models.Message `json:"messages"`
	ImageUrl       string            `json:"image_url"`
	Description    string            `json:"description"`
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	CreatedAt      int64             `json:"created_at"`
	UpdatedAt      int64             `json:"updated_at"`
}

type AddParticipantRequest struct {
	ChatRoomId string   `json:"chat_room_id" binding:"required"`
	UserIds    []string `json:"user_ids" binding:"required"`
}
