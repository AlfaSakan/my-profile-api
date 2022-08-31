package models

type Contact struct {
	UserId   string `json:"user_id" gorm:"not null"`
	FriendId string `json:"friend_id" gorm:"not null"`
	Status   string `json:"status" gorm:"default:active"`
}
