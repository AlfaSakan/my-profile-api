package schemas

type UserRequest struct {
	CountryCode string `json:"country_code" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Name        string `json:"name" binding:"required"`
	ImageUrl    string `json:"image_url"`
	Status      string `json:"status"`
}

type UpdateUserRequest struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	Status      string `json:"status"`
}
