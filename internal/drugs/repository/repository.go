package repository

import (
	"pharmacy/internal/drugs/models"

	"gorm.io/gorm"
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

func (repo *DrugRepository) GetAllDrugs() ([]models.Drug, error) {
	var drugs []models.Drug
	if err := repo.DB.Find(&drugs).Error; err != nil {
		return nil, err
	}
	return drugs, nil
}

func (repo *DrugRepository) GetDrugByID(id uint) (*models.Drug, error) {
	var drug models.Drug
	if err := repo.DB.First(&drug, id).Error; err != nil {
		return nil, err
	}
	return &drug, nil
}