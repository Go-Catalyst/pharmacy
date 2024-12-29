package models


import (
	"gorm.io/gorm"
	// "pharmacy/internal/drugs/models"
)

type Category struct {
	gorm.Model
	// ID          uint   `json:"id" gorm:"primaryKey"`                  
	Name        string `json:"name" gorm:"unique;not null"`
	// Description    string    `json:"description" gorm:"type:text"`
	// CreatedAt      time.Time `json:"created_at"`
	// UpdatedAt      time.Time `json:"updated_at"`
	// IsActive       bool      `json:"is_active" gorm:"default:true"`
	// Slug           string    `json:"slug" gorm:"unique;not null"`
	// Drugs []models.Drug `gorm:"many2many:category_drugs"`
}


