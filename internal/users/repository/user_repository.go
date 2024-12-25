package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pharmacy/internal/users/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db, err := gorm.Open(sqlite.Open("phdb.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetAllUsers() []models.User {
	var users []models.User
	repo.db.Find(&users)
	return users
}

func (repo *UserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) {
	repo.db.Create(user)
}

func (repo *UserRepository) UpdateUser(id string, updatedUser *models.User) (models.User, error) {
	var user models.User
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	repo.db.Save(&user)
	return user, nil
}

func (repo *UserRepository) DeleteUser(id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
