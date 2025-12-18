package service

import (
	"errors"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"

	"github.com/gin-gonic/gin"
)

type ScannerService struct {
	repo *repository.ScannerRepository
}

func NewScannerService(repo *repository.ScannerRepository) *ScannerService {
	return &ScannerService{repo: repo}
}

func getOrgIDFromContext(c *gin.Context) (uint, error) {
	orgIDValue, exists := c.Get("organization_id")
	if !exists {
		return 0, errors.New("organization_id missing in token")
	}

	orgIDFloat, ok := orgIDValue.(float64)
	if !ok {
		return 0, errors.New("invalid organization_id format")
	}

	return uint(orgIDFloat), nil
}

func (s *ScannerService) Create(c *gin.Context, scanner *model.Scanner) error {
	orgID, err := getOrgIDFromContext(c)
	if err != nil {
		return err
	}

	// Enforce ownership
	scanner.OrganizationID = orgID

	return s.repo.Create(c, scanner)
}

func (s *ScannerService) GetAll(c *gin.Context) ([]model.Scanner, error) {
	orgID, err := getOrgIDFromContext(c)
	if err != nil {
		return nil, err
	}

	return s.repo.FindAllByOrg(c, orgID)
}

func (s *ScannerService) GetByID(c *gin.Context, id uint) (*model.Scanner, error) {
	orgID, err := getOrgIDFromContext(c)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(c, id, orgID)
}
func (s *ScannerService) GetByName(c *gin.Context, name string) (*model.Scanner, error) {
	orgID, err := getOrgIDFromContext(c)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByName(c, name, orgID)
}

func (s *ScannerService) Delete(c *gin.Context, id uint) error {
	orgID, err := getOrgIDFromContext(c)
	if err != nil {
		return err
	}

	return s.repo.Delete(c, id, orgID)
}
