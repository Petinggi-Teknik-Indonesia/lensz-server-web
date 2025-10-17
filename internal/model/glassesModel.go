package model

import (
	"gorm.io/gorm"
)

type Drawer struct {
	gorm.Model
	Name    string    `json:"name"`
	Glasses []Glasses `json:"glasses" gorm:"foreignKey:DrawerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Brand struct {
	gorm.Model
	Name    string    `json:"name"`
	Glasses []Glasses `json:"glasses" gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
	gorm.Model
	Name    string    `json:"name"`
	Glasses []Glasses `json:"glasses" gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Glasses struct {
	gorm.Model
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Color       string        `json:"color"`
	Description *string       `json:"description"`
	RFID        *string       `json:"rfid"`
	Status      GlassesStatus `json:"status" gorm:"not null;default:0"`

	DrawerID  uint `json:"-"`
	BrandID   uint `json:"-"`
	CompanyID uint `json:"-"`

	Drawer  Drawer  `json:"drawer"`
	Brand   Brand   `json:"brand"`
	Company Company `json:"company"`

	StatusHistory []StatusHistory `json:"statusHistory" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Enum
type GlassesStatus int

const (
	Tersedia GlassesStatus = iota
	Terjual
	Rusak
	Terpinjam
	Lainnya
)

func (s GlassesStatus) String() string {
	return [...]string{"Tersedia", "Terjual", "Rusak", "Terpinjam", "Lainnya"}[s]
}

type StatusHistory struct {
	gorm.Model
	StatusChange GlassesStatus `json:"statusChange"`
	GlassesID    uint          `json:"glassesId"`
	UserID       uint          `json:"userId"`
	Glasses      Glasses       `json:"glasses"`
	User         User          `json:"user"`
}
