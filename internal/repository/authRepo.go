package repository

import (
	"lensz-server-web/internal/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Role").Preload("Organization").First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) GetRoleByName(name string) (*model.Role, error) {
	var role model.Role
	if err := r.db.First(&role, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
