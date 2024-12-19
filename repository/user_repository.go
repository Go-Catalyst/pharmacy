package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pharmacy/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db, err := gorm.Open("sqlite3", "phdb.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
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
