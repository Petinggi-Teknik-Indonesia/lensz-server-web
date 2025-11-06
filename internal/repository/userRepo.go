package repository

import (
	"context"
	"lensz-server-web/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindUserByEmail(ctx context.Context,email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Role").Preload("Organization").First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindUserByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Role").Preload("Organization").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ---------------- ORGANIZATION ---------------
func (r *UserRepository) CreateOrganization(ctx context.Context, organization  *model.Organization) error {
	return r.db.WithContext(ctx).Create(organization).Error
}

func (r *UserRepository) FindOrganizationByID(ctx context.Context, id uint) (*model.Organization, error) {
	var organization model.Organization
	err := r.db.WithContext(ctx).First(&organization, id).Error
	if err != nil {
		return nil, err
	}
	return &organization, nil
}

func (r *UserRepository) FindAllOrganizations(ctx context.Context) ([]model.Organization, error) {
	var organization []model.Organization
	err := r.db.WithContext(ctx).Find(&organization).Error
	return organization, err
}

func (r *UserRepository) UpdateOrganization(ctx context.Context, organization *model.Organization) error {
	return r.db.WithContext(ctx).Save(organization).Error
}

func (r *UserRepository) DeleteOrganization(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Organization{}, id).Error
}

// ---------------- ROLE ---------------
func (r *UserRepository) CreateRole(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *UserRepository) FindRoleByID(ctx context.Context, id uint) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *UserRepository) FindAllRoles(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.WithContext(ctx).Find(&roles).Error
	return roles, err
}

func (r *UserRepository) UpdateRole(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *UserRepository) DeleteRole(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Role{}, id).Error
}

// ---------------- SCANNER ---------------
func (r *UserRepository) CreateScanner(ctx context.Context, scanner *model.Scanner) error {
	return r.db.WithContext(ctx).Create(scanner).Error
}

func (r *UserRepository) FindScannerByID(ctx context.Context, id uint) (*model.Scanner, error) {
	var scanner model.Scanner
	err := r.db.WithContext(ctx).
		Preload("Organization").
		First(&scanner, id).Error
	if err != nil {
		return nil, err
	}
	return &scanner, nil
}

func (r *UserRepository) FindAllScanners(ctx context.Context) ([]model.Scanner, error) {
	var scanners []model.Scanner
	err := r.db.WithContext(ctx).
		Preload("Organization").
		Find(&scanners).Error
	return scanners, err
}

func (r *UserRepository) UpdateScanner(ctx context.Context, scanner *model.Scanner) error {
	return r.db.WithContext(ctx).Save(scanner).Error
}

func (r *UserRepository) DeleteScanner(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Scanner{}, id).Error
}