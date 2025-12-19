package repository

import (
	"context"
	"lensz-server-web/internal/model"

	"gorm.io/gorm"
)

type ScannerRepository struct {
	db *gorm.DB
}

func NewScannerRepository(db *gorm.DB) *ScannerRepository {
	return &ScannerRepository{db: db}
}

func (r *ScannerRepository) Create(ctx context.Context, scanner *model.Scanner) error {
	return r.db.WithContext(ctx).Create(scanner).Error
}

func (r *ScannerRepository) FindByID(ctx context.Context, id uint, orgID uint) (*model.Scanner, error) {
	var scanner model.Scanner
	err := r.db.WithContext(ctx).
		Where("id = ? AND organization_id = ?", id, orgID).
		First(&scanner).Error
	if err != nil {
		return nil, err
	}
	return &scanner, nil
}

func (r *ScannerRepository) FindByName(ctx context.Context, name string) (*model.Scanner, error) {
	var scanner model.Scanner
	err := r.db.WithContext(ctx).
		Where("device_name = ?", name).
		First(&scanner).Error
	if err != nil {
		return nil, err
	}
	return &scanner, nil
}

func (r *ScannerRepository) FindAllByOrg(ctx context.Context, orgID uint) ([]model.Scanner, error) {
	var scanners []model.Scanner
	err := r.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&scanners).Error
	return scanners, err
}

func (r *ScannerRepository) Delete(ctx context.Context, id uint, orgID uint) error {
	return r.db.WithContext(ctx).
		Where("id = ? AND organization_id = ?", id, orgID).
		Delete(&model.Scanner{}).Error
}
