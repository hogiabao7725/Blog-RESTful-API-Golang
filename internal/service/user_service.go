package service

import (
	"context"
	"errors"
	"fmt"

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
		return nil, fmt.Errorf("service.user.register: %w", err)
	}
	return newUser, nil
}

func (s *userService) Login(ctx context.Context, usernameOrEmail, password string) (*domain.User, error) {

	user, err := s.repo.FindByUsernameOrEmail(ctx, usernameOrEmail)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("service.user.login: %w", err)
	}

	if password != user.Password {
		return nil, domain.ErrInvalidCredentials
	}

	return user, nil
}
