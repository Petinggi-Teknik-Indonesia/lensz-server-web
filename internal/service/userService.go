package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *model.User) error {
	// check duplicate
	existing, _ := s.repo.FindByEmail(user.Email)
	if existing != nil && existing.Email != "" {
		return errors.New("email already registered")
	}

	if err := user.HashPassword(user.Password); err != nil {
		return err
	}

	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !user.CheckPassword(password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("missing JWT_SECRET")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
