package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

type commentService struct {
	repo     domain.CommentRepository
	postRepo domain.PostRepository
}

func NewCommentService(repo domain.CommentRepository, postRepo domain.PostRepository) domain.CommentService {
	return &commentService{
		repo:     repo,
		postRepo: postRepo,
	}
}

func (s *commentService) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	comment.Body = strings.TrimSpace(comment.Body)
	if comment.Body == "" {
		return nil, errorx.NewInvalidInputError("body", "must not be empty")
	}
	if comment.UserID <= 0 {
		return nil, errorx.NewInvalidInputError("user_id", "must be greater than 0")
	}
	if comment.PostID <= 0 {
		return nil, errorx.NewInvalidInputError("post_id", "must be greater than 0")
	}

	if _, err := s.postRepo.FindByID(ctx, comment.PostID); err != nil {
		return nil, fmt.Errorf("service.comment.create.validate_post: %w", err)
	}

	createdComment, err := s.repo.Create(ctx, comment)
	if err != nil {
		return nil, fmt.Errorf("service.comment.create: %w", err)
	}

	return createdComment, nil
}

func (s *commentService) FindByID(ctx context.Context, id int64) (*domain.Comment, error) {
	comment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service.comment.find_by_id: %w", err)
	}
	return comment, nil
}

func (s *commentService) FindByPostID(ctx context.Context, postID int64) ([]*domain.Comment, error) {
	if _, err := s.postRepo.FindByID(ctx, postID); err != nil {
		return nil, fmt.Errorf("service.comment.find_by_post_id.validate_post: %w", err)
	}

	comments, err := s.repo.FindByPostID(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("service.comment.find_by_post_id: %w", err)
	}
	return comments, nil
}

func (s *commentService) Update(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	comment.Body = strings.TrimSpace(comment.Body)
	if comment.Body == "" {
		return nil, errorx.NewInvalidInputError("body", "must not be empty")
	}

	updatedComment, err := s.repo.Update(ctx, comment)
	if err != nil {
		return nil, fmt.Errorf("service.comment.update: %w", err)
	}
	return updatedComment, nil
}

func (s *commentService) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("service.comment.delete: %w", err)
	}
	return nil
}
