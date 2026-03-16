package request

type PostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	UserID      int64  `json:"user_id"`
	CategoryID  int64  `json:"category_id"`
}
