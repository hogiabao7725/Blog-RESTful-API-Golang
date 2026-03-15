package errorx

import "errors"

// -- Helper functions to check error types --

func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}

func IsForbidden(err error) bool {
	return errors.Is(err, ErrForbidden)
}

func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}

// -- Helper functions for custom error types --
// -- Is Not Found Error --
func IsNotFound(err error) bool {
	var e *NotFoundError
	return errors.As(err, &e)
}

func NewNotFoundError(resource, field string, value any) *NotFoundError {
	return &NotFoundError{
		Resource: resource,
		Field:    field,
		Value:    value,
	}
}

// -- Is Already Exists Error --
func IsAlreadyExists(err error) bool {
	var e *AlreadyExistsError
	return errors.As(err, &e)
}

func NewAlreadyExistsError(resource, field string, value any) *AlreadyExistsError {
	return &AlreadyExistsError{
		Resource: resource,
		Field:    field,
		Value:    value,
	}
}

// -- Is Invalid Input Error --
func IsInvalidInput(err error) bool {
	var e *InvalidInputError
	return errors.As(err, &e)
}

func NewInvalidInputError(field, message string) *InvalidInputError {
	return &InvalidInputError{
		Field:   field,
		Message: message,
	}
}

// -- Is Invalid Credentials Error --
func IsInvalidCredentials(err error) bool {
	var e *InvalidCredentialsError
	return errors.As(err, &e)
}

func NewInvalidCredentialsError(msg ...string) *InvalidCredentialsError {
	if len(msg) > 0 {
		return &InvalidCredentialsError{Message: msg[0]}
	}
	return &InvalidCredentialsError{}
}
