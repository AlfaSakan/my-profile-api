package schemas

type SessionRequest struct {
	UserAgent   string `json:"user_agent"`
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
