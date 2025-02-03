package models

import (
	"gorm.io/gorm"
	"pharmacy/internal/drugs/models"
)

type Category struct {
	gorm.Model
	Name        string        `json:"name" gorm:"unique;not null"`
	Description string        `json:"description" gorm:"type:text"`
	IsActive    bool          `gorm:"default:true"`
	Slug        string        `json:"slug" gorm:"unique;not null"`
	Drugs       []models.Drug `gorm:"many2many:category_drugs"`
}
