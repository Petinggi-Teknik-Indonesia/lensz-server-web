package model

import (
	"gorm.io/gorm"
)

type Scanner struct {
	gorm.Model
	Name string

}
type Ownership struct{
	gorm.Model
	OrganizationID uint
	Organization Organization
	Scanner []Scanner
}