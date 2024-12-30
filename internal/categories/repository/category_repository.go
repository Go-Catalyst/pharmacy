package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pharmacy/internal/categories/models"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	db, err := gorm.Open("sqlite3", "phdb.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Category{})
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAllCategories() []models.Category {
	var categories []models.Category
	repo.db.Find(&categories)
	return categories
}

func (repo *CategoryRepository) GetCategoryByID(id string) (models.Category, error) {
	var category models.Category
	if err := repo.db.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (repo *CategoryRepository) CreateCategory(category *models.Category) {
	repo.db.Create(&category)
}

func (repo *CategoryRepository) UpdateCategory(id string, updatedCategory models.Category) (models.Category, error) {
	var category models.Category
	if err := repo.db.First(&category, id).Error; err != nil {
		return category, err
	}
	repo.db.Model(&category).Updates(updatedCategory)
	return category, nil
}

func (repo *CategoryRepository) DeleteCategory(id string) error {
	if err := repo.db.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
