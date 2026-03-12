package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
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
				return nil, fmt.Errorf("email '%s' already registered: %w", user.Email, domain.ErrAlreadyExists)
			case "users_username_key":
				return nil, fmt.Errorf("username '%s' is taken: %w", user.Username, domain.ErrAlreadyExists)
			default:
				return nil, fmt.Errorf("conflict on %s: %w", pgErr.Constraint, domain.ErrAlreadyExists)
			}
		}
		// Always use %w to retain the original error trace stack from the database for debugging (logging).
		return nil, fmt.Errorf("db.user.create: %w", err)
	}

	return user, nil
}

func (r *userRepository) Get(ctx context.Context, id int64) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
