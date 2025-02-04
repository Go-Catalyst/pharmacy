package repository

import (
	"pharmacy/internal/categories/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
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
