package model

import (
	"gorm.io/gorm"
)

// Each scanner belongs to one organization
type Scanner struct {
	gorm.Model
	Name           string       `json:"name"`
	OrganizationID uint         `json:"organizationId"` // foreign key
	Organization   Organization `json:"organization" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
