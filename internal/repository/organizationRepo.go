package repository

import (
	"context"
	"lensz-server-web/internal/model"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) CreateOrganization(ctx context.Context, org *model.Organization) error {
	return r.db.WithContext(ctx).Create(org).Error
}

func (r *OrganizationRepository) FindAllOrganizations(ctx context.Context) ([]model.Organization, error) {
	var orgs []model.Organization
	err := r.db.WithContext(ctx).Find(&orgs).Error
	return orgs, err
}

func (r *OrganizationRepository) FindOrganizationByID(ctx context.Context, id uint) (*model.Organization, error) {
	var org model.Organization
	if err := r.db.WithContext(ctx).Preload("User").Preload("Scanners").First(&org, id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (r *OrganizationRepository) UpdateOrganization(ctx context.Context, org *model.Organization) error {
	return r.db.WithContext(ctx).Save(org).Error
}

func (r *OrganizationRepository) DeleteOrganization(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Organization{}, id).Error
}
