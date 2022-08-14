package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/schemas"
)

type IUserService interface {
	FindUserById(userId uint) (models.User, error)
	CreateUser(schemas.UserRequest) (models.User, error)
	UpdateUser(schemas.UpdateUserRequest, uint) (models.User, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (userService *UserService) FindUserById(userId uint) (models.User, error) {
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

func (userService *UserService) UpdateUser(userRequest schemas.UpdateUserRequest, userId uint) (models.User, error) {
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
