package service

import (
	"context"
	"fmt"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
)

type postService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) domain.PostService {
	return &postService{
		repo: repo,
	}
}

func (s *postService) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	newPost, err := s.repo.Create(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("service.post.create: %w", err)
	}
	return newPost, nil
}

func (s *postService) FindByID(ctx context.Context, id int64) (*domain.Post, error) {
	post, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service.post.find_by_id: %w", err)
	}
	return post, nil
}

func (s *postService) FindAll(ctx context.Context) ([]*domain.Post, error) {
	posts, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.post.find_all: %w", err)
	}
	return posts, nil
}

func (s *postService) FindByCategoryID(ctx context.Context, categoryID int64) ([]*domain.Post, error) {
	posts, err := s.repo.FindByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("service.post.find_by_category_id: %w", err)
	}
	return posts, nil
}

func (s *postService) Update(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	updatedPost, err := s.repo.Update(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("service.post.update: %w", err)
	}
	return updatedPost, nil
}

func (s *postService) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("service.post.delete: %w", err)
	}
	return nil
}
