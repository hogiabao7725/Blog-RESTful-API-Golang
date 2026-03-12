package service

import (
	"context"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user.RoleID == 0 {
		user.RoleID = 2 // Default to "user" role if not specified.
	}
	newUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
