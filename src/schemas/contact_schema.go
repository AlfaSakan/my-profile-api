package schemas

type ContactRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	FriendId string `json:"friend_id" binding:"required"`
	Status   string `json:"status"`
}
