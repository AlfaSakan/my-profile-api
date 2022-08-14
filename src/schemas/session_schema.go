package schemas

type SessionRequest struct {
	UserId    int    `json:"user_id" binding:"required"`
	UserAgent string `json:"user_agent"`
}
