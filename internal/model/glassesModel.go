package model

import (
	"gorm.io/gorm"
)

type Drawer struct {
	gorm.Model
	Name    string
	Glasses []Glasses `gorm:"foreignKey:DrawerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Brand struct {
	gorm.Model
	Name    string
	Glasses []Glasses `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
	gorm.Model
	Name    string
	Glasses []Glasses `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Glasses struct {
	gorm.Model
	Name        string
	Type        string
	Color       string
	Description *string
	RFID *string
	Status      GlassesStatus `gorm:"not null"`

	// Foreign Keys (nullable because of SET NULL)
	DrawerID  *uint
	BrandID   *uint
	CompanyID *uint

	// Relations
	Drawer  Drawer
	Brand   Brand
	Company Company
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
