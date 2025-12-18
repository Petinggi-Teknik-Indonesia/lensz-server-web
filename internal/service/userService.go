package service

import (
	"context"
	"errors"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Helpers
func (s *UserService) generateAccessToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":           user.ID,
		"role":         user.RoleID,
		"email":        user.Email,
		"organization": user.OrganizationID,
		"type":         "access",
		"exp":          time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *UserService) generateRefreshToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.ID,
		"type": "refresh",
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}


type UserService struct {
	repo      *repository.UserRepository
	jwtSecret string
}

func NewUserService(repo *repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{repo: repo, jwtSecret: jwtSecret}
}

func (s *UserService) Register(ctx context.Context, u *model.User) error {
	existing, _ := s.repo.FindUserByEmail(ctx, u.Email)
	if existing != nil {
		return errors.New("user already exists")
	}

	if err := u.HashPassword(u.Password); err != nil {
		return err
	}
	u.VerifiedStatus = false // pending
	return s.repo.CreateUser(ctx, u)
}

// Admin registers directly (backdoor)
func (s *UserService) AdminRegister(ctx context.Context, u *model.User) error {
	if err := u.HashPassword(u.Password); err != nil {
		return err
	}
	u.VerifiedStatus = true // verified
	return s.repo.CreateUser(ctx, u)
}

// Admin verifies a user
func (s *UserService) VerifyUser(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.VerifiedStatus = true
	updatedUser, err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// Admin rejects a user
func (s *UserService) CancelUser(ctx context.Context, email string) error {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}
	return s.repo.DeleteUser(ctx, user)
}

// Login (for both users and admins)
func (s *UserService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil || !user.CheckPassword(password) {
		return "", "", errors.New("invalid email or password")
	}

	if !user.VerifiedStatus {
		return "", "", errors.New("account not yet verified by admin")
	}

	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "refresh" {
		return "", errors.New("invalid refresh token type")
	}

	userID := uint(claims["id"].(float64))

	user, err := s.repo.FindUserByID(ctx, userID)
	if err != nil {
		return "", errors.New("user not found")
	}

	return s.generateAccessToken(user)
}



// Fetch all unverified users
func (s *UserService) GetAllUnverified(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAllUnverified(ctx)
}

// Fetch unverified users in specific organization
func (s *UserService) GetUnverifiedByOrg(ctx context.Context, orgID uint) ([]model.User, error) {
	org, err := s.repo.FindOrganizationByID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindUserUnverified(ctx, org)
}

// Fetch verified users in specific organization
func (s *UserService) GetVerifiedByOrg(ctx context.Context, orgID uint) ([]model.User, error) {
	org, err := s.repo.FindOrganizationByID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindUserVerified(ctx, org)
}

// Fetch all verified users
func (s *UserService) GetAllVerified(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAllVerified(ctx)
}
