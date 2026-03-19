package request

import (
	"strings"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

type CreateCommentRequest struct {
	Body string `json:"body"`
}

type UpdateCommentRequest struct {
	Body string `json:"body"`
}

func (r *CreateCommentRequest) Normalize() {
	r.Body = strings.TrimSpace(r.Body)
}

func (r *CreateCommentRequest) Validate() error {
	if len(r.Body) == 0 {
		return errorx.NewInvalidInputError("body", "is required")
	}
	return nil
}

func (r *UpdateCommentRequest) Normalize() {
	r.Body = strings.TrimSpace(r.Body)
}

func (r *UpdateCommentRequest) Validate() error {
	if len(r.Body) == 0 {
		return errorx.NewInvalidInputError("body", "is required")
	}
	return nil
}
