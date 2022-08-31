package repositories

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindUserById(userId string) (models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(user models.User, userId string) (models.User, error)
	FindUser(*models.User) (*models.User, error)
	SearchNameUser(name string) (*[]models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository *UserRepository) FindUserById(userId string) (models.User, error) {
	var user models.User

	err := repository.db.Where(&models.User{UserId: userId}).Find(&user).Error

	return user, err
}

func (repository *UserRepository) CreateUser(user models.User) (models.User, error) {
	err := repository.db.Create(&user).Error

	return user, err
}

func (repository *UserRepository) UpdateUser(user models.User, userId string) (models.User, error) {
	err := repository.db.Where(&models.User{UserId: userId}).Updates(&user).Error

	return user, err
}

func (repository *UserRepository) FindUser(user *models.User) (*models.User, error) {
	foundUser := &models.User{}

	err := repository.db.Find(foundUser, user).Error

	return foundUser, err
}

func (r *UserRepository) SearchNameUser(name string) (*[]models.User, error) {
	users := &[]models.User{}
	err := r.db.Where("name like ?", "%"+name+"%").Find(users).Error

	return users, err
}
