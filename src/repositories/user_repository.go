package repositories

import (
	"myProfileApi/src/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindUserById(uint) (models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(models.User, uint) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository *UserRepository) FindUserById(userId uint) (models.User, error) {
	var user models.User

	err := repository.db.Debug().Find(&user, userId).Error

	return user, err
}

func (repository *UserRepository) CreateUser(user models.User) (models.User, error) {
	err := repository.db.Debug().Create(&user).Error

	return user, err
}

func (repository *UserRepository) UpdateUser(user models.User, userId uint) (models.User, error) {
	err := repository.db.Where(&models.User{UserId: uint(userId)}).Updates(&user).Error

	return user, err
}
