package models

type Session struct {
	SessionId uint   `json:"session_id" gorm:"primaryKey"`
	Valid     bool   `json:"valid" gorm:"default:true"`
	UserAgent string `json:"user_agent" gorm:"size:200; not null"`
	UserId    uint   `json:"user_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
