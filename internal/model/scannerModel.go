package model

import (
	"gorm.io/gorm"
)

type Scanner struct {
	gorm.Model
	Name string
}

//  {"Name":"value"} 
//  {"name":"value"} `json:"name"`

type Ownership struct{
	gorm.Model
	OrganizationID uint 
	Organization Organization
	Scanner []Scanner `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
