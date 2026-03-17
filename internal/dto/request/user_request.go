package request

import (
	"net/mail"
	"strings"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterUserRequest) Normalize() {
	r.Name = strings.Join(strings.Fields(r.Name), " ")
	r.Username = strings.TrimSpace(r.Username)
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))
}

func (r *RegisterUserRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errorx.NewInvalidInputError("name", "is required")
	}
	if len(strings.TrimSpace(r.Username)) < 3 {
		return errorx.NewInvalidInputError("username", "must be at least 3 characters")
	}
	if strings.TrimSpace(r.Email) == "" {
		return errorx.NewInvalidInputError("email", "is required")
	}
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errorx.NewInvalidInputError("email", "is not a valid email address")
	}
	if len(r.Password) < 8 {
		return errorx.NewInvalidInputError("password", "must be at least 8 characters")
	}
	return nil
}

type LoginUserRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

func (r *LoginUserRequest) Normalize() {
	r.UsernameOrEmail = strings.TrimSpace(r.UsernameOrEmail)
}

func (r *LoginUserRequest) Validate() error {
	if strings.TrimSpace(r.UsernameOrEmail) == "" {
		return errorx.NewInvalidInputError("username or email", "is required")
	}
	if len(r.Password) == 0 {
		return errorx.NewInvalidInputError("password", "is required")
	}
	return nil
}
