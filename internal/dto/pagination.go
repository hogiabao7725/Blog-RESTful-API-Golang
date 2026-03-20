package dto

// PaginationRequest holds pagination parameters from query string
type PaginationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// PaginationMeta holds pagination metadata for response
type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

// GetOffset calculates database offset from page and limit
func (p *PaginationRequest) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

// NewPaginationMeta creates pagination metadata
func NewPaginationMeta(page, limit int, total int64) *PaginationMeta {
	totalPages := (total + int64(limit) - 1) / int64(limit)
	return &PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    int64(page) < totalPages,
		HasPrev:    page > 1,
	}
}
