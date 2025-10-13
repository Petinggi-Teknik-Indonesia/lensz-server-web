package repository

import (
	"context"
	"lensz-server-web/internal/model"

	"gorm.io/gorm"
)

type GlassesRepository struct {
	db *gorm.DB
}

func NewGlassesRepository(db *gorm.DB) *GlassesRepository {
	return &GlassesRepository{db: db}
}

// ------------------- GLASSES -------------------

// Create Glasses
func (r *GlassesRepository) CreateGlasses(ctx context.Context, glasses *model.Glasses) error {
	return r.db.WithContext(ctx).Create(glasses).Error
}

// Get Glasses by ID (with relations)
func (r *GlassesRepository) FindGlassesByID(ctx context.Context, id uint) (*model.Glasses, error) {
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
func (r *GlassesRepository) FindAllGlasses(ctx context.Context) ([]model.Glasses, error) {
	var glasses []model.Glasses
	err := r.db.WithContext(ctx).
		Preload("Drawer").
		Preload("Brand").
		Preload("Company").
		Find(&glasses).Error
	return glasses, err
}

func (r *GlassesRepository) FindAllGlassesPartial(ctx context.Context) ([]model.GlassesPartialResponse, error) {
	var result []model.GlassesPartialResponse

	err := r.db.WithContext(ctx).
		Table("glasses").
		Select("glasses.id, glasses.name, glasses.color, glasses.status, brands.name as brand, drawers.name as drawer").
		Joins("LEFT JOIN brands ON brands.id = glasses.brand_id").
		Joins("LEFT JOIN drawers ON drawers.id = glasses.drawer_id").
		Scan(&result).Error

	return result, err
}

func (r *GlassesRepository) FindGlassesSimplifiedByID(ctx context.Context, id uint) (*model.GlassesSingleResponse, error) {
	var result model.GlassesSingleResponse

	err := r.db.WithContext(ctx).
		Table("glasses").
		Select(`
			glasses.id,
			glasses.name,
			glasses.type,
			glasses.color,
			glasses.description,
			glasses.rf_id,
			glasses.status,
			drawers.name as drawer,
			brands.name as brand,
			companies.name as company,
			glasses.created_at,
			glasses.updated_at
		`).
		Joins("LEFT JOIN drawers ON drawers.id = glasses.drawer_id").
		Joins("LEFT JOIN brands ON brands.id = glasses.brand_id").
		Joins("LEFT JOIN companies ON companies.id = glasses.company_id").
		Where("glasses.id = ?", id).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}




// Update Glasses
func (r *GlassesRepository) UpdateGlasses(ctx context.Context, glasses *model.Glasses) error {
	return r.db.WithContext(ctx).Save(glasses).Error
}

// Delete Glasses
func (r *GlassesRepository) DeleteGlasses(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Glasses{}, id).Error
}

// Find Glasses by Status
func (r *GlassesRepository) FindGlassesByStatus(ctx context.Context, status model.GlassesStatus) ([]model.Glasses, error) {
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
func (r *GlassesRepository) FindGlassesByBrand(ctx context.Context, brandID uint) ([]model.Glasses, error) {
	var glasses []model.Glasses
	err := r.db.WithContext(ctx).
		Where("brand_id = ?", brandID).
		Preload("Drawer").
		Preload("Brand").
		Preload("Company").
		Find(&glasses).Error
	return glasses, err
}

// ------------------- DRAWER -------------------

func (r *GlassesRepository) CreateDrawer(ctx context.Context, drawer *model.Drawer) error {
	return r.db.WithContext(ctx).Create(drawer).Error
}

func (r *GlassesRepository) FindDrawerByID(ctx context.Context, id uint) (*model.Drawer, error) {
	var drawer model.Drawer
	err := r.db.WithContext(ctx).First(&drawer, id).Error
	if err != nil {
		return nil, err
	}
	return &drawer, nil
}

func (r *GlassesRepository) FindAllDrawers(ctx context.Context) ([]model.Drawer, error) {
	var drawers []model.Drawer
	err := r.db.WithContext(ctx).Find(&drawers).Error
	return drawers, err
}

// ------------------- BRAND -------------------

func (r *GlassesRepository) CreateBrand(ctx context.Context, brand *model.Brand) error {
	return r.db.WithContext(ctx).Create(brand).Error
}

func (r *GlassesRepository) FindBrandByID(ctx context.Context, id uint) (*model.Brand, error) {
	var brand model.Brand
	err := r.db.WithContext(ctx).First(&brand, id).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (r *GlassesRepository) FindAllBrands(ctx context.Context) ([]model.Brand, error) {
	var brands []model.Brand
	err := r.db.WithContext(ctx).Find(&brands).Error
	return brands, err
}

// ------------------- COMPANY -------------------

func (r *GlassesRepository) CreateCompany(ctx context.Context, company *model.Company) error {
	return r.db.WithContext(ctx).Create(company).Error
}

func (r *GlassesRepository) FindCompanyByID(ctx context.Context, id uint) (*model.Company, error) {
	var company model.Company
	err := r.db.WithContext(ctx).First(&company, id).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *GlassesRepository) FindAllCompanies(ctx context.Context) ([]model.Company, error) {
	var companies []model.Company
	err := r.db.WithContext(ctx).Find(&companies).Error
	return companies, err
}
