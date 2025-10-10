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


