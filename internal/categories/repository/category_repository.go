package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"pharmacy/internal/categories/models"
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
	repo.db.Create(&category)
}

func (repo *CategoryRepository) GetAllCategories() []models.Category {
	var categories []models.Category
	repo.db.Find(&categories)
	return categories
}

// GetCategoryByID todo: test it!
func (repo *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category *models.Category
	log.Print(id)
	if err := repo.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *CategoryRepository) UpdateCategory(id uint, updatedCategory models.Category) (*models.Category, error) {
	var category *models.Category
	if err := repo.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	repo.db.Model(&category).Updates(updatedCategory)
	return category, nil
}

func (repo *CategoryRepository) DeleteCategory(id uint) error {
	err := repo.db.Delete(&models.Category{}, id).Error
	return err
}
