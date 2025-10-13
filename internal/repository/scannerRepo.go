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

// Create pending RFID if not exists
func (r *ScannerRepository) CreatePendingRFID(ctx context.Context, p *model.PendingRFID) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *ScannerRepository) FindPendingByRFID(ctx context.Context, rfid string) (*model.PendingRFID, error) {
	var p model.PendingRFID
	err := r.db.WithContext(ctx).Where("rfid = ?", rfid).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ScannerRepository) MarkPendingRegistered(ctx context.Context, rfid string) error {
	return r.db.WithContext(ctx).Model(&model.PendingRFID{}).
		Where("rfid = ?", rfid).
		Update("registered", true).Error
}
