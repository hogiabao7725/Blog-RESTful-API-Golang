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

// PaginatedComments holds paginated comments with total count
type PaginatedComments struct {
	Comments []*Comment
	Total    int64
}

type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	FindByID(ctx context.Context, id int64) (*Comment, error)
	FindByPostID(ctx context.Context, postID int64) ([]*Comment, error)
	FindByPostIDPaginated(ctx context.Context, postID int64, offset, limit int) (*PaginatedComments, error)
	Update(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int64) error
}

type CommentService interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	FindByID(ctx context.Context, id int64) (*Comment, error)
	FindByPostID(ctx context.Context, postID int64) ([]*Comment, error)
	FindByPostIDPaginated(ctx context.Context, postID int64, offset, limit int) (*PaginatedComments, error)
	Update(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int64) error
}
