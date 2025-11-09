package service

import (
	"context"
	"errors"
	"fmt"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}
func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, u *model.User) error {
	existing, _ := s.repo.FindUserByEmail(ctx, u.Email)
	if existing != nil {
		return errors.New("user already exists")
	}

	if err := u.HashPassword(u.Password); err != nil {
		return err
	}
	u.VerifiedStatus = 0 // pending
	return s.repo.CreateUser(ctx, u)
}

// Admin registers directly (backdoor)
func (s *UserService) AdminRegister(ctx context.Context, u *model.User) error {
	if err := u.HashPassword(u.Password); err != nil {
		return err
	}
	u.VerifiedStatus = 1 // verified
	return s.repo.CreateUser(ctx, u)
}

// Admin verifies a user
func (s *UserService) VerifyUser(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.VerifiedStatus = 1
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
func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid email or password")
	}

	if user.VerifiedStatus == 0 {
		return nil, errors.New("account not yet verified by admin")
	}

	return user, nil
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
package service

import (
	"context"
	"errors"
	"fmt"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}
func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, u *model.User) error {
	existing, _ := s.repo.FindUserByEmail(ctx, u.Email)
	if existing != nil {
		return errors.New("user already exists")
	}

	if err := u.HashPassword(u.Password); err != nil {
		return err
	}
	u.VerifiedStatus = 0 // pending
	return s.repo.CreateUser(ctx, u)
}

// Admin registers directly (backdoor)
func (s *UserService) AdminRegister(ctx context.Context, u *model.User) error {
	if err := u.HashPassword(u.Password); err != nil {
		return err
	}
	u.VerifiedStatus = 1 // verified
	return s.repo.CreateUser(ctx, u)
}

// Admin verifies a user
func (s *UserService) VerifyUser(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.VerifiedStatus = 1
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
func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid email or password")
	}

	if user.VerifiedStatus == 0 {
		return nil, errors.New("account not yet verified by admin")
	}

	return user, nil
}

// Fetch all unverified users
func (s *UserService) GetAllUnverified(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAllUnverified(ctx)
}
// Fetch all verified users
func (s *UserService) GetAllVerified(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAllVerified(ctx)
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