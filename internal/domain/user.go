package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*User, error)
}

type UserService interface {
	Register(ctx context.Context, user *User) (*User, error)
	Login(ctx context.Context, username, password string) (*User, error)
}
