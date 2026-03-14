package domain

import "context"

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	FindByID(ctx context.Context, id int64) (*Category, error)
	FindAll(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) (*Category, error)
	Delete(ctx context.Context, id int64) error
}

type CategoryService interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	FindByID(ctx context.Context, id int64) (*Category, error)
	FindAll(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) (*Category, error)
	Delete(ctx context.Context, id int64) error
}
