package model

import (
	"gorm.io/gorm"
)

// Each scanner belongs to one organization
type Scanner struct {
	gorm.Model
	DeviceName     string       `json:"deviceName" gorm:"unique;not null"`
	OrganizationID uint         `json:"organizationId"` // foreign key
	Organization   Organization `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
