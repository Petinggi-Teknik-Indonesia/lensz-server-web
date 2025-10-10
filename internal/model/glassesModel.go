package model

import (
	"gorm.io/gorm"
)

type Drawer struct {
	gorm.Model
	Name    string     
	Glasses []Glasses  `gorm:"foreignKey:DrawerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Brand struct {
	gorm.Model
	Name    string     
	Glasses []Glasses  `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
	gorm.Model
	Name    string     
	Glasses []Glasses  `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Glasses struct {
	gorm.Model
	Name        string         
	Type        string         
	Color       string         
	Description *string        
	RFID        *string        
	Status      GlassesStatus  `gorm:"not null;default:0"`

	DrawerID  uint    
	BrandID   uint    
	CompanyID uint    

	Drawer  Drawer  
	Brand   Brand   
	Company Company 

	StatusHistory []StatusHistory  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

type StatusHistory struct{
	gorm.Model
	StatusChange GlassesStatus
	GlassesID uint
	UserID uint

	Glasses Glasses 
	User User
}

