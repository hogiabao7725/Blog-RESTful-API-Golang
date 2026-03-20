package utils

import (
	"net/http"
	"strconv"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

const (
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100
)

// ParsePagination extracts and validates pagination parameters from query string
// Query parameters: page (default 1), limit (default 10, max 100)
func ParsePagination(r *http.Request) (*dto.PaginationRequest, error) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := DefaultPage
	limit := DefaultLimit

	// Parse page parameter
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil || p < 1 {
			return nil, errorx.NewInvalidInputError("page", "must be a positive integer")
		}
		page = p
	}

	// Parse limit parameter
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil || l < 1 {
			return nil, errorx.NewInvalidInputError("limit", "must be a positive integer")
		}
		if l > MaxLimit {
			l = MaxLimit
		}
		limit = l
	}

	return &dto.PaginationRequest{
		Page:  page,
		Limit: limit,
	}, nil
}
