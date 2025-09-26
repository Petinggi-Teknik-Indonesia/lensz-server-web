package service

import (
	"errors"
	"lensz-server-web/internal/domain/enums"
	"lensz-server-web/internal/domain/models"
	"lensz-server-web/internal/repository"
)

type ItemService interface {
	GetAll() ([]models.Item, error)
	GetByID(id string) (*models.Item, error)
	Create(item *models.Item) error
	Update(id string, updated *models.Item) error
	Delete(id string) error
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) GetAll() ([]models.Item, error) {
	return s.repo.FindAll()
}

func (s *itemService) GetByID(id string) (*models.Item, error) {
	return s.repo.FindByID(id)
}

func (s *itemService) Create(item *models.Item) error {
	// validate status
	switch item.Status {
	case enums.StatusTersedia, enums.StatusTerjual, enums.StatusRusak, enums.StatusTerpinjam, enums.StatusLainnya:
		return s.repo.Create(item)
	default:
		return errors.New("invalid status")
	}
}

func (s *itemService) Update(id string, updated *models.Item) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// update fields
	existing.Drawer = updated.Drawer
	existing.Color = updated.Color
	existing.Type = updated.Type
	existing.Brand = updated.Brand
	existing.Company = updated.Company
	existing.Status = updated.Status

	return s.repo.Update(existing)
}

func (s *itemService) Delete(id string) error {
	return s.repo.Delete(id)
}
