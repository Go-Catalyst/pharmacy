package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `gorm:"unique;not null" json:"name"`
	Drugs []Drug `gorm:"many2many:category_drugs"`
}
