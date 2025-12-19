package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type DrawerScanSession struct {
	ID            string
	DrawerID      uint
	ExpectedTotal int

	ScannedRFIDs map[string]bool
	Counted      int

	Status    string // running | completed | cancelled
	StartedAt time.Time
}

type DrawerScanService struct {
	sessions      map[string]*DrawerScanSession
	activeSession *DrawerScanSession
	mu            sync.Mutex
	repo          *repository.GlassesRepository
}

type ScanResult struct {
	Missing   []model.Glasses
	Mislabels []model.Glasses
	Scanned   int
	Expected  int
}

func NewDrawerScanService(repo *repository.GlassesRepository) *DrawerScanService {
	return &DrawerScanService{
		sessions: make(map[string]*DrawerScanSession),
		repo:     repo,
	}
}

func (s *DrawerScanService) StartSession(
	ctx context.Context,
	drawerID uint,
	sessionID string,
) (*DrawerScanSession, error) {

	glasses, err := s.repo.FindGlassesByDrawerID(ctx, drawerID)
	if err != nil {
		return nil, err
	}

	session := &DrawerScanSession{
		ID:            sessionID,
		DrawerID:      drawerID,
		ExpectedTotal: len(glasses),
		ScannedRFIDs:  make(map[string]bool),
		Status:        "running",
		StartedAt:     time.Now(),
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.activeSession = session
	s.sessions[sessionID] = session

	return session, nil
}

func (s *DrawerScanService) HandleRFID(
	ctx context.Context,
	sessionID string,
	rfid string,
) (int, int, error) {

	s.mu.Lock()
	session, ok := s.sessions[sessionID]
	s.mu.Unlock()

	if !ok || session.Status != "running" {
		return 0, 0, errors.New("no active session")
	}

	if session.ScannedRFIDs[rfid] {
		return session.Counted, session.ExpectedTotal, nil
	}

	glass, err := s.repo.FindGlassesByRFID(ctx, rfid)
	if err != nil || glass.DrawerID != session.DrawerID {
		return session.Counted, session.ExpectedTotal, nil
	}

	session.ScannedRFIDs[rfid] = true
	session.Counted++

	return session.Counted, session.ExpectedTotal, nil
}

func (s *DrawerScanService) StopSession(
	ctx context.Context,
	sessionID string,
) (*ScanResult, error) {

	s.mu.Lock()
	session, ok := s.sessions[sessionID]
	if ok {
		session.Status = "completed"
		if s.activeSession != nil && s.activeSession.ID == sessionID {
			s.activeSession = nil
		}
	}
	s.mu.Unlock()

	if !ok {
		return nil, errors.New("session not found")
	}

	all, err := s.repo.FindGlassesByDrawerID(ctx, session.DrawerID)
	if err != nil {
		return nil, err
	}

	var missing []model.Glasses
	var mislabels []model.Glasses

	for _, g := range all {
		if g.RFID == nil {
			continue
		}

		if session.ScannedRFIDs[*g.RFID] {
			if g.Status != model.Tersedia {
				mislabels = append(mislabels, g)
			}
		} else {
			missing = append(missing, g)
		}
	}

	return &ScanResult{
		Missing:   missing,
		Mislabels: mislabels,
		Scanned:   session.Counted,
		Expected:  session.ExpectedTotal,
	}, nil
}

type ScanProgress struct {
	Counted  int
	Expected int
}

func (s *DrawerScanService) TryHandleRFID(
	ctx context.Context,
	rfid string,
) (*ScanProgress, *model.HardwareScanResponse) {

	s.mu.Lock()
	session := s.activeSession
	s.mu.Unlock()

	// ❌ No active session
	if session == nil || session.Status != "running" {
		return nil, &model.HardwareScanResponse{
			Status:  "inactive",
			Message: "No active drawer scan",
		}
	}

	// 🔁 Duplicate scan
	if session.ScannedRFIDs[rfid] {
		return &ScanProgress{
				Counted:  session.Counted,
				Expected: session.ExpectedTotal,
			}, &model.HardwareScanResponse{
				Status:  "ignored",
				Message: "Duplicate scan ignored",
				Counted: session.Counted,
				Expected: session.ExpectedTotal,
			}
	}

	counted, expected, err :=
		s.HandleRFID(ctx, session.ID, rfid)

	// ❌ RFID invalid / wrong drawer / DB error
	if err != nil {
		return nil, &model.HardwareScanResponse{
			Status:  "error",
			Message: "RFID not recognized for this drawer",
		}
	}

	// ✅ Successful scan
	return &ScanProgress{
			Counted:  counted,
			Expected: expected,
		}, &model.HardwareScanResponse{
			Status:  "ok",
			Message: fmt.Sprintf("%d / %d glasses scanned", counted, expected),
			Counted: counted,
			Expected: expected,
		}
}

