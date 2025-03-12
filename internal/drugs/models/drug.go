package models

import (
	"gorm.io/gorm"
	"time"
)


type Drug struct {
	gorm.Model
	Name       string `json:"name"`
	Category   string `json:"category"`
	Expiration time.Time `gorm:"not null" json:"exp"`
	Doses      int    `json:"doses"`
}
