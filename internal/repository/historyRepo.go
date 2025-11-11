package repository

import (
	"context"
	"lensz-server-web/internal/model"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (r *HistoryRepository) CreateHistory(ctx context.Context, history *model.StatusHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

func (r *HistoryRepository) FindHistoryByGlassesID(ctx context.Context, glassesID uint) ([]model.StatusHistoryResponse, error) {
	var histories []model.StatusHistoryResponse

	err := r.db.WithContext(ctx).
		Table("status_histories").
		Select("id", "status_change", "glasses_id", "user_id", "created_at").
		Where("glasses_id = ?", glassesID).
		Order("created_at DESC").
		Find(&histories).Error

	if err != nil {
		return nil, err
	}

	return histories, nil
}
