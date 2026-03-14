package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func ParseIDFromURL(r *http.Request) (int64, error) {
	idStr := r.PathValue("id")

	if idStr == "" {
		return 0, fmt.Errorf("id is required")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("utils.parse_id_from_url: %w", err)
	}

	return id, nil
}
