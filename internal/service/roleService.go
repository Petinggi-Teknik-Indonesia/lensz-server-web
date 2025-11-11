package service

import (
	"context"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(ctx context.Context, role *model.Role) error {
	return s.repo.CreateRole(ctx, role)
}

func (s *RoleService) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	return s.repo.FindAllRoles(ctx)
}

func (s *RoleService) GetRoleByID(ctx context.Context, id uint) (*model.Role, error) {
	return s.repo.FindRoleByID(ctx, id)
}

func (s *RoleService) UpdateRole(ctx context.Context, role *model.Role) error {
	return s.repo.UpdateRole(ctx, role)
}

func (s *RoleService) DeleteRole(ctx context.Context, id uint) error {
	return s.repo.DeleteRole(ctx, id)
}
