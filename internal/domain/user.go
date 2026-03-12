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
	Get(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}

type UserService interface {
	Register(ctx context.Context, user *User) (*User, error)
}
