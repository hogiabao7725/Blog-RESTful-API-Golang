package domain

import (
	"context"
	"time"
)

type Post struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	UserID      int64     `json:"user_id"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PaginatedPosts holds paginated posts with total count
type PaginatedPosts struct {
	Posts []*Post
	Total int64
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	FindByID(ctx context.Context, id int64) (*Post, error)
	FindAll(ctx context.Context) ([]*Post, error)
	FindAllPaginated(ctx context.Context, offset, limit int) (*PaginatedPosts, error)
	FindByCategoryID(ctx context.Context, categoryID int64) ([]*Post, error)
	FindByCategoryIDPaginated(ctx context.Context, categoryID int64, offset, limit int) (*PaginatedPosts, error)
	Search(ctx context.Context, query string) ([]*Post, error)
	SearchPaginated(ctx context.Context, query string, offset, limit int) (*PaginatedPosts, error)
	Update(ctx context.Context, post *Post) (*Post, error)
	Delete(ctx context.Context, id int64) error
}

type PostService interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	FindByID(ctx context.Context, id int64) (*Post, error)
	FindAll(ctx context.Context) ([]*Post, error)
	FindAllPaginated(ctx context.Context, offset, limit int) (*PaginatedPosts, error)
	FindByCategoryID(ctx context.Context, categoryID int64) ([]*Post, error)
	FindByCategoryIDPaginated(ctx context.Context, categoryID int64, offset, limit int) (*PaginatedPosts, error)
	Search(ctx context.Context, query string) ([]*Post, error)
	SearchPaginated(ctx context.Context, query string, offset, limit int) (*PaginatedPosts, error)
	Update(ctx context.Context, post *Post) (*Post, error)
	Delete(ctx context.Context, id int64) error
}
