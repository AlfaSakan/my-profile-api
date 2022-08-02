package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/schemas"
)

type IUserService interface {
	// FindAllUser() ([]models.User, error)
	FindUserById(userId int) (models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(models.User, int) (models.User, error)
}

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (userService *UserService) FindUserById(userId int) (models.User, error) {
	user, err := userService.userRepository.FindUserById(userId)
	return user, err
}

func (userService *UserService) CreateUser(userRequest schemas.UserRequest) (models.User, error) {
	user := models.User{
		Name:        userRequest.Name,
		CountryCode: userRequest.CountryCode,
		PhoneNumber: userRequest.PhoneNumber,
		ImageUrl:    userRequest.ImageUrl,
		Status:      userRequest.Status,
	}

	user, err := userService.userRepository.CreateUser(user)
	return user, err
}

func (userService *UserService) UpdateUser(userRequest schemas.UpdateUserRequest, userId int) (models.User, error) {
	user := models.User{
		Name:        userRequest.Name,
		CountryCode: userRequest.CountryCode,
		PhoneNumber: userRequest.PhoneNumber,
		ImageUrl:    userRequest.ImageUrl,
		Status:      userRequest.Status,
	}

	user, err := userService.userRepository.UpdateUser(user, userId)
	return user, err
}
