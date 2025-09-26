package repository

import (
	"lensz-server-web/internal/domain/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() ([]models.Item, error)
	FindByID(id string) (*models.Item, error)
	Create(item *models.Item) error
	Update(item *models.Item) error
	Delete(id string) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) FindAll() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) FindByID(id string) (*models.Item, error) {
	var item models.Item
	err := r.db.First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) Update(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id string) error {
	return r.db.Delete(&models.Item{}, "id = ?", id).Error
}
