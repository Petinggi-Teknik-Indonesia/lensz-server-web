package service

import (
	"context"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type OrganizationService struct {
	repo *repository.OrganizationRepository
}

func NewOrganizationService(repo *repository.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, org *model.Organization) error {
	return s.repo.CreateOrganization(ctx, org)
}

func (s *OrganizationService) GetAllOrganizations(ctx context.Context) ([]model.Organization, error) {
	return s.repo.FindAllOrganizations(ctx)
}

func (s *OrganizationService) GetOrganizationByID(ctx context.Context, id uint) (*model.Organization, error) {
	return s.repo.FindOrganizationByID(ctx, id)
}

func (s *OrganizationService) UpdateOrganization(ctx context.Context, org *model.Organization) error {
	return s.repo.UpdateOrganization(ctx, org)
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, id uint) error {
	return s.repo.DeleteOrganization(ctx, id)
}
