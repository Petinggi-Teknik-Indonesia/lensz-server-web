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
	Status      GlassesStatus `json:"status" gorm:"type:varchar(20);not null;default:'Tersedia'"`

	DrawerID  uint `json:"drawerId"`
	BrandID   uint `json:"brandId"`
	CompanyID uint `json:"companyId"`
	OrganizationID uint         `json:"organizationId"`

	Drawer  Drawer  `json:"drawer"`
	Brand   Brand   `json:"brand"`
	Company Company `json:"company"`
    Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	StatusHistory []StatusHistory `json:"statusHistory" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	
}

// Enum
type GlassesStatus string

const (
	Tersedia  GlassesStatus = "Tersedia"
	Terjual   GlassesStatus = "Terjual"
	Rusak     GlassesStatus = "Rusak"
	Terpinjam GlassesStatus = "Terpinjam"
	Lainnya   GlassesStatus = "Lainnya"
)

// func (s GlassesStatus) String() string {
// 	return [...]string{"Tersedia", "Terjual", "Rusak", "Terpinjam", "Lainnya"}[s]
// }

type StatusHistory struct {
	gorm.Model
	StatusChange GlassesStatus `json:"statusChange" gorm:"type:varchar(20);not null"`
	GlassesID    uint          `json:"glassesId"`
	UserID       uint          `json:"userId"`
	OrganizationID uint		   `json:"organizationId"`
	Organization Organization
	Glasses      Glasses       `json:"glasses"`
	User         User          `json:"user"`
}
