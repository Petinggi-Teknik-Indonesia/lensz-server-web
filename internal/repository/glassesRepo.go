package repository

import (
	"context"
	"lensz-server-web/internal/model"
	"gorm.io/gorm"
)

type GlassesRepository struct {
	db *gorm.DB
}

func NewGlassesRepository(db *gorm.DB) *GlassesRepository{
	return &GlassesRepository{db: db}
}

// Create Glasses
func (r *GlassesRepository) Create(ctx context.Context, glasses *model.Glasses) error {
	return r.db.WithContext(ctx).Create(glasses).Error
}


// Get Glasses by ID (with relations)
func (r *GlassesRepository) FindByID(ctx context.Context, id uint) (*model.Glasses, error) {
	var glasses model.Glasses
	err := r.db.WithContext(ctx).
		Preload("Drawer").
		Preload("Brand").
		Preload("Company").
		First(&glasses, id).Error
	if err != nil {
		return nil, err
	}
	return &glasses, nil
}


// Get All Glasses
func (r *GlassesRepository) FindAll(ctx context.Context) ([]model.Glasses, error) {
	var glasses []model.Glasses
	err := r.db.WithContext(ctx).
		Preload("Drawer").
		Preload("Brand").
		Preload("Company").
		Find(&glasses).Error
	return glasses, err
}

// Update Glasses
func (r *GlassesRepository) Update(ctx context.Context, glasses *model.Glasses) error {
	return r.db.WithContext(ctx).Save(glasses).Error
}

// Delete Glasses
func (r *GlassesRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Glasses{}, id).Error
}

// Find Glasses by Status
func (r *GlassesRepository) FindByStatus(ctx context.Context, status model.GlassesStatus) ([]model.Glasses, error) {
	var glasses []model.Glasses
	err := r.db.WithContext(ctx).
		Where("status = ?", status).
		Preload("Drawer").
		Preload("Brand").
		Preload("Company").
		Find(&glasses).Error
	return glasses, err
}

// Find Glasses by Brand
func (r *GlassesRepository) FindByBrand(ctx context.Context, brandID uint) ([]model.Glasses, error) {
	var glasses []model.Glasses
	err := r.db.WithContext(ctx).
		Where("brand_id = ?", brandID).
		Preload("Drawer").
		Preload("Brand").
		Preload("Company").
		Find(&glasses).Error
	return glasses, err
}