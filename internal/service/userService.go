package service

import (
	"errors"
	"lensz-server-web/internal/domain/models"
	"lensz-server-web/internal/repository"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Register(user *models.User) error {
	// Hash password here (bcrypt)
	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	// Compare hashed passwords
	if password != user.Password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
