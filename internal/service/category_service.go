package service

import (
	"context"
	"fmt"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
)

type categoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) domain.CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	newCategory, err := s.repo.Create(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("service.category.create: %w", err)
	}
	return newCategory, nil
}

func (s *categoryService) FindByID(ctx context.Context, id int64) (*domain.Category, error) {
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service.category.find_by_id: %w", err)
	}
	return category, nil
}
func (s *categoryService) FindAll(ctx context.Context) ([]*domain.Category, error) {
	categories, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.category.find_all: %w", err)
	}
	return categories, nil
}

func (s *categoryService) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	_, err := s.repo.FindByID(ctx, category.ID)
	if err != nil {
		return nil, fmt.Errorf("service.category.update: %w", err)
	}
	updatedCategory, err := s.repo.Update(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("service.category.update: %w", err)
	}
	return updatedCategory, nil
}

func (s *categoryService) Delete(ctx context.Context, id int64) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("service.category.delete: %w", err)
	}
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("service.category.delete: %w", err)
	}
	return nil
}
