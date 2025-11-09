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

func (s *UserService) Register(ctx contect.Context)