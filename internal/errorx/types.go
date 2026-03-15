package errorx

import (
	"errors"
	"fmt"
)

// Sensinel errors - use for default error because no more information is needed.
var (
	ErrUnauthorized = errors.New("unauthorized")

	ErrForbidden = errors.New("forbidden")

	ErrInternal = errors.New("internal error")
)

// -- Custom Domain Errors --

// -- Not Found Error --

type NotFoundError struct {
	Resource string
	Field    string
	Value    any
}

func (e *NotFoundError) Error() string {
	if e.Field != "" && e.Value != nil {
		return fmt.Sprintf("%s with %s '%v' not found", e.Resource, e.Field, e.Value)
	}
	if e.Resource != "" {
		return fmt.Sprintf("%s not found", e.Resource)
	}
	return "resource not found"
}

// -- Already Exists Error --

type AlreadyExistsError struct {
	Resource string
	Field    string
	Value    any
}

func (e *AlreadyExistsError) Error() string {
	if e.Field != "" && e.Value != nil {
		return fmt.Sprintf("%s with %s '%v' already exists", e.Resource, e.Field, e.Value)
	}
	if e.Resource != "" {
		return fmt.Sprintf("%s already exists", e.Resource)
	}
	return "resource already exists"
}

// -- Invalid Input Error --

type InvalidInputError struct {
	Field   string
	Message string
}

func (e *InvalidInputError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s", e.Field, e.Message)
	}
	return e.Message
}

// -- Invalid Credentials Error --

type InvalidCredentialsError struct {
	Message string
}

func (e *InvalidCredentialsError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "invalid credentials"
}
