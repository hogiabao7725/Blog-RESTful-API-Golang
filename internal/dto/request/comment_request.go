package request

type CreateCommentRequest struct {
	Body   string `json:"body"`
	UserID int64  `json:"user_id"`
}

type UpdateCommentRequest struct {
	Body string `json:"body"`
}
