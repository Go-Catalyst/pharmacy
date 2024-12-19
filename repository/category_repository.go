package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pharmacy/models"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	db, err := gorm.Open(sqlite.Open("phdb.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.Category{})
	if err != nil {
		panic("failed to migrate database" + err.Error())
	}
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) CreateCategory(category *models.Category) {
	repo.db.Create(category)
}
