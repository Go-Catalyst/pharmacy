package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email       string     `gorm:"unique;not null" json:"email"`
	Username    string     `gorm:"unique;not null" json:"username"`
	Password    string     `gorm:"not null" json:"password"`
	FirstName   string     `gorm:"not null" json:"firstName"`
	LastName    string     `gorm:"not null" json:"lastName"`
	PhoneNumber string     `gorm:"unique;not null" json:"phoneNumber"`
	Role        string     `gorm:"not null" json:"role"`
	IsActive    bool       `gorm:"default:true" json:"isActive"`
	LastLoginAt *time.Time `gorm:"default:null" json:"lastLoginAt"`
}
