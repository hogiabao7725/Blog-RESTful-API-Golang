package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	Body      string    `json:"body"`
	UserID    int64     `json:"user_id"`
	PostID    int64     `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	FindByID(ctx context.Context, id int64) (*Comment, error)
	FindByPostID(ctx context.Context, postID int64) ([]*Comment, error)
	Update(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int64) error
}

type CommentService interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	FindByID(ctx context.Context, id int64) (*Comment, error)
	FindByPostID(ctx context.Context, postID int64) ([]*Comment, error)
	Update(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int64) error
}
