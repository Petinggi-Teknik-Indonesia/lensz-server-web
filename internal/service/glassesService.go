package service

import (
	"context"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type GlassesService struct {
	repo *repository.GlassesRepository
}

func NewGlassesService(repo *repository.GlassesRepository) *GlassesService {
	return &GlassesService{repo: repo}
}

func (s *GlassesService) CreateGlasses(ctx context.Context, g *model.Glasses) error {
	return s.repo.Create(ctx, g)
}

func (s *GlassesService) GetGlassesByID(ctx context.Context, id uint) (*model.Glasses, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *GlassesService) GetAllGlasses(ctx context.Context) ([]model.Glasses, error) {
	return s.repo.FindAll(ctx)
}

func (s *GlassesService) UpdateGlasses(ctx context.Context, g *model.Glasses) error {
	return s.repo.Update(ctx, g)
}

func (s *GlassesService) DeleteGlasses(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *GlassesService) GetGlassesByStatus(ctx context.Context, status model.GlassesStatus) ([]model.Glasses, error) {
	return s.repo.FindByStatus(ctx, status)
}

func (s *GlassesService) GetGlassesByBrand(ctx context.Context, brandID uint) ([]model.Glasses, error) {
	return s.repo.FindByBrand(ctx, brandID)
}
