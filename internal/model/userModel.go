package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique;not null"`
	Users []User `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
    gorm.Model
    Name           string `json:"name"`
    Email          string `json:"email" gorm:"uniqueIndex;not null"`
    Phone          string `json:"phone"`
    Password       string `json:"-" gorm:"not null"`
    VerifiedStatus bool   `json:"verifiedStatus" gorm:"default:false"`
    RoleID         uint   `json:"roleId"`
    Role           Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

    OrganizationID uint         `json:"organizationId"`
    Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}


type Organization struct {
	gorm.Model
	Name     string    `json:"name"`
	Scanners []Scanner `json:"scanners" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User     []User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// HashPassword generates a bcrypt hash
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword verifies a bcrypt hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
