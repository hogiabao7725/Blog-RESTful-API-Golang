package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
	"github.com/lib/pq"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
			INSERT INTO users (name, username, email, password, role_id)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, created_at
	`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Username,
		user.Email, user.Password, user.RoleID).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			// Optimization: Return the error with context so the Handler knows how to report it to the User.
			switch pgErr.Constraint {
			// case "table_column_key":
			// 	return nil, fmt.Errorf("value '%s' already exists: %w", value, domain.ErrAlreadyExists)
			case "users_email_key":
				return nil, errorx.NewAlreadyExistsError("user", "email", user.Email)
			case "users_username_key":
				return nil, errorx.NewAlreadyExistsError("user", "username", user.Username)
			default:
				return nil, errorx.NewAlreadyExistsError("user", "constraint", pgErr.Constraint)
			}
		}
		// Always use %w to retain the original error trace stack from the database for debugging (logging).
		return nil, fmt.Errorf("db.user.create: %w", err)
	}

	return user, nil
}

func (r *userRepository) FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*domain.User, error) {
	query := `
		SELECT id, name, username, email, password, role_id, created_at
		FROM users
		WHERE username = $1 OR email = $1
	`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, usernameOrEmail).Scan(
		&user.ID, &user.Name, &user.Username, &user.Email,
		&user.Password, &user.RoleID, &user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// This will be used by logging, not use by service or handler, so we can provide more context for debugging.
			return nil, errorx.NewNotFoundError("user", "username OR email", usernameOrEmail)
		}
		return nil, fmt.Errorf("db.user.findByUsernameOrEmail: %w", err)
	}
	return user, nil
}
