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

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	query := `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
		RETURNING id
	`
	err := r.db.QueryRowContext(ctx, query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.Constraint {
			case "categories_name_unique":
				return nil, errorx.NewAlreadyExistsError("category", "name", category.Name)
			default:
				return nil, errorx.NewAlreadyExistsError("category", "constraint", pgErr.Constraint)
			}
		}
		return nil, fmt.Errorf("db.category.create: %w", err)
	}
	return category, nil
}

func (r *categoryRepository) FindByID(ctx context.Context, id int64) (*domain.Category, error) {
	query := `
		SELECT id, name, description
		FROM categories
		WHERE id = $1
	`
	category := &domain.Category{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NewNotFoundError("category", "id", id)
		}
		return nil, fmt.Errorf("db.category.find_by_id: %w", err)
	}
	return category, nil
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]*domain.Category, error) {
	query := `
		SELECT id, name, description
		FROM categories
		ORDER BY id ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("db.category.find_all: %w", err)
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		category := &domain.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, fmt.Errorf("db.category.find_all: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	query := `
		UPDATE categories
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id
	`
	err := r.db.QueryRowContext(ctx, query, category.Name, category.Description, category.ID).Scan(&category.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NewNotFoundError("category", "id", category.ID)
		}
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.Constraint {
			case "categories_name_unique":
				return nil, errorx.NewAlreadyExistsError("category", "name", category.Name)
			default:
				return nil, errorx.NewAlreadyExistsError("category", "constraint", pgErr.Constraint)
			}
		}
		return nil, fmt.Errorf("db.category.update: %w", err)
	}
	return category, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM categories
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("db.category.delete: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("db.category.delete: %w", err)
	}
	if rowsAffected == 0 {
		return errorx.NewNotFoundError("category", "id", id)
	}
	return nil
}
