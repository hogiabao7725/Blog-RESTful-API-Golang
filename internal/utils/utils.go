package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

func ParseIDFromURL(r *http.Request) (int64, error) {
	return ParsePathID(r, "id")
}

func ParsePathID(r *http.Request, paramName string) (int64, error) {
	value := strings.TrimSpace(r.PathValue(paramName))
	if value == "" {
		return 0, errorx.NewInvalidInputError(paramName, "missing identifier in URL")
	}

	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errorx.NewInvalidInputError(paramName, fmt.Sprintf("invalid ID format '%s'", value))
	}

	if id <= 0 {
		return 0, errorx.NewInvalidInputError(paramName, "id must be greater than 0")
	}

	return id, nil
}
