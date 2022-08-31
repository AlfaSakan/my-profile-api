package services

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"
)

type IUserService interface {
	FindUserById(userId string) (models.User, error)
	CreateUser(schemas.UserRequest) (models.User, error)
	UpdateUser(schemas.UpdateUserRequest, string) (models.User, error)
	FindUser(*schemas.UserRequest) (*models.User, error)
	SearchUserByName(name string) (*[]models.User, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (userService *UserService) FindUserById(userId string) (models.User, error) {
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
		UserId:      utils.GenerateId(),
	}

	user, err := userService.userRepository.CreateUser(user)
	return user, err
}

func (userService *UserService) UpdateUser(userRequest schemas.UpdateUserRequest, userId string) (models.User, error) {
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

func (userService *UserService) FindUser(userRequest *schemas.UserRequest) (*models.User, error) {
	user := &models.User{
		Name:        userRequest.Name,
		CountryCode: userRequest.CountryCode,
		PhoneNumber: userRequest.PhoneNumber,
	}

	foundUser, err := userService.userRepository.FindUser(user)

	return foundUser, err
}

func (us *UserService) SearchUserByName(name string) (*[]models.User, error) {
	users, err := us.userRepository.SearchNameUser(name)
	return users, err
}
