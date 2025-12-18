// often called usecase
package service

import (
	"context"
	"errors"
	"fmt"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type GlassesService struct {
	repo        *repository.GlassesRepository
	historyRepo *repository.HistoryRepository
}

func NewGlassesService(repo *repository.GlassesRepository, historyRepo *repository.HistoryRepository) *GlassesService {
	return &GlassesService{repo: repo, historyRepo: historyRepo}
}

func (s *GlassesService) CreateGlasses(ctx context.Context, g *model.Glasses, userID uint) error {
	// ---- HANDLE BRAND ----
	fmt.Println(g);
	if g.Brand.ID != 0 {
		// check existing brand
		brand, err := s.repo.FindBrandByID(ctx, g.Brand.ID)
		if err != nil {
			return errors.New("brand not found")
		}
		g.BrandID = brand.ID
	} else if g.Brand.Name != "" {
		// create new brand
		if err := s.repo.CreateBrand(ctx, &g.Brand); err != nil {
			return err
		}
		g.BrandID = g.Brand.ID
	} else {
		return errors.New("brand is required (id or name)")
	}

	// ---- HANDLE COMPANY ----
	if g.Company.ID != 0 {
		company, err := s.repo.FindCompanyByID(ctx, g.Company.ID)
		if err != nil {
			return errors.New("company not found")
		}
		g.CompanyID = company.ID
	} else if g.Company.Name != "" {
		if err := s.repo.CreateCompany(ctx, &g.Company); err != nil {
			return err
		}
		g.CompanyID = g.Company.ID
	} else {
		return errors.New("company is required (id or name)")
	}

	// ---- HANDLE DRAWER ----
	if g.Drawer.ID != 0 {
		drawer, err := s.repo.FindDrawerByID(ctx, g.Drawer.ID)
		if err != nil {
			return errors.New("drawer not found")
		}
		g.DrawerID = drawer.ID
	} else if g.Drawer.Name != "" {
		if err := s.repo.CreateDrawer(ctx, &g.Drawer); err != nil {
			return err
		}
		g.DrawerID = g.Drawer.ID
	} else {
		return errors.New("drawer is required (id or name)")
	}
	// ---- Create Glasses ----
	if err := s.repo.CreateGlasses(ctx, g); err != nil {
		return err
	}
	

	// logging
	history := &model.StatusHistory{
		StatusChange: g.Status,
		GlassesID:    g.ID,
		UserID:       userID, 
	}
	return s.historyRepo.CreateHistory(ctx, history)
}

func (s *GlassesService) GetGlassesByID(ctx context.Context, id uint) (*model.Glasses, error) {
	return s.repo.FindGlassesByID(ctx, id)
}

func (s *GlassesService) GetAllGlasses(ctx context.Context) ([]model.Glasses, error) {
	return s.repo.FindAllGlasses(ctx)
}

func (s *GlassesService) GetAllGlassesPartial(ctx context.Context) ([]model.GlassesPartialResponse, error) {
	return s.repo.FindAllGlassesPartial(ctx)
}

func (s *GlassesService) GetGlassesSimplifiedByID(ctx context.Context, id uint) (*model.GlassesSingleResponse, error) {
	return s.repo.FindGlassesSimplifiedByID(ctx, id)
}




func (s *GlassesService) UpdateGlasses(ctx context.Context, g *model.Glasses, userID uint) error {
	existing, err := s.repo.FindGlassesByID(ctx, g.ID)
	if err != nil {
		return errors.New("glasses not found")
	}

	statusChanged := existing.Status != g.Status
	if err := s.repo.UpdateGlasses(ctx, g); err != nil {
		return err
	}

	if statusChanged {
		history := &model.StatusHistory{
			StatusChange: g.Status,
			GlassesID:    g.ID,
			UserID:       userID,
		}
		return s.historyRepo.CreateHistory(ctx, history)
	}

	return nil
}

func (s *GlassesService) UpdateGlassesStatusByRFID(ctx context.Context, rfid string, newStatus model.GlassesStatus, userID uint) error {
	// Find glasses by RFID
	glasses, err := s.repo.FindGlassesByRFID(ctx, rfid)
	if err != nil {
		return errors.New("glasses not found for given RFID")
	}

	// Update status only
	glasses.Status = newStatus
	if err := s.repo.UpdateGlasses(ctx, glasses); err != nil {
		return err
	}

	// Log the change
	history := &model.StatusHistory{
		StatusChange: newStatus,
		GlassesID:    glasses.ID,
		UserID:       userID,
	}
	return s.historyRepo.CreateHistory(ctx, history)
}

func (s *GlassesService) DeleteGlasses(ctx context.Context, id uint) error {
	return s.repo.DeleteGlasses(ctx, id)
}

func (s *GlassesService) GetGlassesByStatus(ctx context.Context, status model.GlassesStatus) ([]model.Glasses, error) {
	return s.repo.FindGlassesByStatus(ctx, status)
}

func (s *GlassesService) GetGlassesByBrand(ctx context.Context, brandID uint) ([]model.Glasses, error) {
	return s.repo.FindGlassesByBrand(ctx, brandID)
}
