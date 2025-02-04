package repository

import (
	"gorm.io/gorm"
	"pharmacy/internal/drugs/models"
)

type DrugRepository struct {
	DB *gorm.DB
}

func NewDrugRepository(db *gorm.DB) *DrugRepository {
	return &DrugRepository{DB: db}
}

func (repo *DrugRepository) CreateDrug(drug *models.Drug) error {
	return repo.DB.Create(drug).Error
}
