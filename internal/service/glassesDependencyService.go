package service

import (
	"context"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type GlassesDependencyService struct {
	repo *repository.GlassesRepository
}

func NewGlassesDependencyService(repo *repository.GlassesRepository) *GlassesDependencyService {
	return &GlassesDependencyService{repo: repo}
}

// -------------------- DRAWER --------------------
func (s *GlassesDependencyService) CreateDrawer(ctx context.Context, d *model.Drawer) error {
	return s.repo.CreateDrawer(ctx, d)
}

func (s *GlassesDependencyService) GetDrawerByID(ctx context.Context, id uint) (*model.Drawer, error) {
	return s.repo.FindDrawerByID(ctx, id)
}

func (s *GlassesDependencyService) GetAllDrawers(ctx context.Context) ([]model.Drawer, error) {
	return s.repo.FindAllDrawers(ctx)
}

func (s *GlassesDependencyService) UpdateDrawer(ctx context.Context, d *model.Drawer) error {
	return s.repo.UpdateDrawer(ctx, d)
}

func (s *GlassesDependencyService) DeleteDrawer(ctx context.Context, id uint) error {
	return s.repo.DeleteDrawer(ctx, id)
}

// -------------------- BRAND --------------------
func (s *GlassesDependencyService) CreateBrand(ctx context.Context, b *model.Brand) error {
	return s.repo.CreateBrand(ctx, b)
}

func (s *GlassesDependencyService) GetBrandByID(ctx context.Context, id uint) (*model.Brand, error) {
	return s.repo.FindBrandByID(ctx, id)
}

func (s *GlassesDependencyService) GetAllBrands(ctx context.Context) ([]model.Brand, error) {
	return s.repo.FindAllBrands(ctx)
}

func (s *GlassesDependencyService) UpdateBrand(ctx context.Context, b *model.Brand) error {
	return s.repo.UpdateBrand(ctx, b)
}

func (s *GlassesDependencyService) DeleteBrand(ctx context.Context, id uint) error {
	return s.repo.DeleteBrand(ctx, id)
}

// -------------------- COMPANY --------------------
func (s *GlassesDependencyService) CreateCompany(ctx context.Context, c *model.Company) error {
	return s.repo.CreateCompany(ctx, c)
}

func (s *GlassesDependencyService) GetCompanyByID(ctx context.Context, id uint) (*model.Company, error) {
	return s.repo.FindCompanyByID(ctx, id)
}

func (s *GlassesDependencyService) GetAllCompanies(ctx context.Context) ([]model.Company, error) {
	return s.repo.FindAllCompanies(ctx)
}

func (s *GlassesDependencyService) UpdateCompany(ctx context.Context, c *model.Company) error {
	return s.repo.UpdateCompany(ctx, c)
}

func (s *GlassesDependencyService) DeleteCompany(ctx context.Context, id uint) error {
	return s.repo.DeleteCompany(ctx, id)
}
