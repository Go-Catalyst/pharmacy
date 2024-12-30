package models

import (
	"gorm.io/gorm"
	"pharmacy/internal/drugs/models"
)

type Category struct {
	gorm.Model
	Name        string        `gorm:"unique;not null" json:"name"`
	Description string        `json:"description"`
	IsActive    bool          `gorm:"default:true"`
	Drugs       []models.Drug `gorm:"many2many:category_drugs"`
}
