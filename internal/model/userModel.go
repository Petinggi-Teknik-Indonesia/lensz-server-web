package model

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Role struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Phone    string 
	Password string `json:"-" gorm:"not null"`
	RoleID   uint
	Role     Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Organization struct {
	gorm.Model
	Name string 
}

type OrganizationMembers struct {
	gorm.Model
	OrganizationID uint 
	UserID uint 

	User []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Organization []Organization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
