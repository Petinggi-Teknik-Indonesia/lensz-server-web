package model

import (
	"gorm.io/gorm"
)

// Each scanner belongs to one organization
type Scanner struct {
	gorm.Model
	Name           string
	OrganizationID uint          // foreign key
	Organization   Organization  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PendingRFID struct {
	gorm.Model
	RFID       string    `gorm:"uniqueIndex;not null" json:"rfid"`
	ScannerID  uint      `json:"scannerId"`
	Scanner    Scanner   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"scanner"`
	Registered bool      `gorm:"default:false" json:"registered"`
}
