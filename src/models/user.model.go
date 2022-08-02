package models

type User struct {
	UserId      uint   `json:"user_id" gorm:"primaryKey"`
	CountryCode string `json:"country_code" gorm:"size:5;not null"`
	PhoneNumber string `json:"phone_number" gorm:"size:20;not null"`
	Name        string `json:"name" gorm:"size:100;not null"`
	ImageUrl    string `json:"image_url" gorm:"size:200"`
	Status      string `json:"status" gorm:"size:20"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
