package request

import (
	"strings"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

type PostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	UserID      int64  `json:"user_id"`
	CategoryID  int64  `json:"category_id"`
}

func (r *PostRequest) Normalize() {
	r.Title = strings.TrimSpace(r.Title)
	r.Description = strings.TrimSpace(r.Description)
	r.Content = strings.TrimSpace(r.Content)
}

func (r *PostRequest) Validate() error {
	if len(r.Title) == 0 {
		return errorx.NewInvalidInputError("title", "is required")
	}
	if len(r.Title) > 100 {
		return errorx.NewInvalidInputError("title", "must be less than 100 characters")
	}
	if len(r.Description) > 250 {
		return errorx.NewInvalidInputError("description", "must be less than 250 characters")
	}
	if len(r.Content) == 0 {
		return errorx.NewInvalidInputError("content", "is required")
	}
	if r.UserID <= 0 {
		return errorx.NewInvalidInputError("user_id", "must be a positive integer")
	}
	if r.CategoryID <= 0 {
		return errorx.NewInvalidInputError("category_id", "must be a positive integer")
	}
	return nil
}
