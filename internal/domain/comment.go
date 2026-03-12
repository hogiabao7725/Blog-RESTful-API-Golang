package domain

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	Body      string    `json:"body"`
	UserID    int64     `json:"user_id"`
	PostID    int64     `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
