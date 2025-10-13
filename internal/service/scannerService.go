package service

import (
	"context"
	"errors"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type ScannerService struct {
	repo *repository.ScannerRepository
	// used to update a glasses record when registering RFID
	glassesRepo *repository.GlassesRepository
}

func NewScannerService(r *repository.ScannerRepository, gRepo *repository.GlassesRepository) *ScannerService {
	return &ScannerService{repo: r, glassesRepo: gRepo}
}

// Called when a device scans an RFID
// returns the created/existing PendingRFID
func (s *ScannerService) CreatePendingRFID(ctx context.Context, p *model.PendingRFID) (*model.PendingRFID, error) {
	// if exists return existing (avoid duplicate)
	existing, err := s.repo.FindPendingByRFID(ctx, p.RFID)
	if err == nil && existing != nil {
		return existing, nil
	}
	if err := s.repo.CreatePendingRFID(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

// Register a glasses to RFID: sets Glasses.RFID and marks pending as registered
func (s *ScannerService) RegisterGlassesToRFID(ctx context.Context, rfid string, glassesID uint) error {
	// get glasses
	glasses, err := s.glassesRepo.FindGlassesByID(ctx, glassesID)
	if err != nil {
		return errors.New("glasses not found")
	}

	// set the RFID field
	glasses.RFID = &rfid
	if err := s.glassesRepo.UpdateGlasses(ctx, glasses); err != nil {
		return err
	}

	// mark pending as registered (if exists)
	_ = s.repo.MarkPendingRegistered(ctx, rfid)
	return nil
}
