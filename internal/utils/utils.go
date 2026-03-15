package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

func ParseIDFromURL(r *http.Request) (int64, error) {
	idStr := r.PathValue("id")

	if idStr == "" {
		return 0, errorx.NewInvalidInputError("id", "missing identifier in URL")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, errorx.NewInvalidInputError("id", fmt.Sprintf("invalid ID format '%s'", idStr))
	}

	return id, nil
}
