package models

import "gorm.io/gorm"

type Drug struct {
	gorm.Model
	Name string
}
