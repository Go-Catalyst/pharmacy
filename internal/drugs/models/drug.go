package models

//import "gorm.io/gorm"

type Drug struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	Expiration string `json:"expiration"`
	Doses      int    `json:"doses"`
}
