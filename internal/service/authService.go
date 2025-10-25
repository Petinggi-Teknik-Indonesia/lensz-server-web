package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type AuthService struct {
	repo      *repository.AuthRepository
	jwtSecret []byte
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (s *AuthService) Signup(name, email, phone, password, roleName string, orgID uint) error {
	role, err := s.repo.GetRoleByName(roleName)
	if err != nil {
		return errors.New("invalid role")
	}

	user := &model.User{
		Name:           name,
		Email:          email,
		Phone:          phone,
		RoleID:         role.ID,
		OrganizationID: orgID,
	}

	if err := user.HashPassword(password); err != nil {
		return err
	}

	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !user.CheckPassword(password) {
		return "", errors.New("invalid email or password")
	}

	claims := jwt.MapClaims{
		"user_id":        user.ID,
		"email":          user.Email,
		"role":           user.Role.Name,
		"organizationId": user.OrganizationID,
		"exp":            time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
