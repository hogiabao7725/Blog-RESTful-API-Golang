package errorx

import (
	"errors"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/response"
)

func WriteError(w http.ResponseWriter, statusCode int, errMsg string) {
	response.WriteJSON(w, statusCode, response.Response{
		Success: false,
		Error:   errMsg,
	})
}

func WriteDomainError(w http.ResponseWriter, err error) {
	// log.Printf("WriteDomainError received: %T - %v", err, err)
	statusCode := http.StatusInternalServerError
	message := "internal server error"

	var (
		notFoundErr      *NotFoundError
		alreadyExistsErr *AlreadyExistsError
		invalidInputErr  *InvalidInputError
		invalidCredErr   *InvalidCredentialsError
	)

	switch {
	case errors.As(err, &notFoundErr):
		statusCode = http.StatusNotFound
		message = notFoundErr.Error()
	case errors.As(err, &alreadyExistsErr):
		statusCode = http.StatusConflict
		message = alreadyExistsErr.Error()
	case errors.As(err, &invalidInputErr):
		statusCode = http.StatusBadRequest
		message = invalidInputErr.Error()
	case errors.As(err, &invalidCredErr):
		statusCode = http.StatusUnauthorized
		message = invalidCredErr.Error()
	case errors.Is(err, ErrUnauthorized):
		statusCode = http.StatusUnauthorized
		message = ErrUnauthorized.Error()
	case errors.Is(err, ErrForbidden):
		statusCode = http.StatusForbidden
		message = ErrForbidden.Error()
	case errors.Is(err, ErrInternal):
		statusCode = http.StatusInternalServerError
		message = ErrInternal.Error()
	}

	WriteError(w, statusCode, message)
}
