package response

// PaginatedResponse is the standard response format for paginated data
type PaginatedResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}
