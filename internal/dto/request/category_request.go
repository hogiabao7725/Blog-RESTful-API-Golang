package request

import (
	"strings"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CategoryRequest) Normalize() {
	r.Name = strings.Join(strings.Fields(r.Name), " ")
	r.Description = strings.TrimSpace(r.Description)
}

func (r *CategoryRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errorx.NewInvalidInputError("name", "is required")
	}
	if len(r.Name) > 100 {
		return errorx.NewInvalidInputError("name", "must be less than 100 characters")
	}
	if len(r.Description) > 255 {
		return errorx.NewInvalidInputError("description", "must be less than 255 characters")
	}
	return nil
}
