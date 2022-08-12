package services

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/schemas"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repositories.IUserRepositoryMock{Mock: mock.Mock{}}
var userService = NewUserService(userRepository)

const findUserById = "FindUserById"
const createUser = "CreateUser"
const updateUser = "UpdateUser"

var user = models.User{
	UserId:      1,
	CountryCode: "+62",
	PhoneNumber: "8123456789",
	Name:        "Test",
	ImageUrl:    "http://test.com",
	Status:      "available",
	CreatedAt:   123,
	UpdatedAt:   123,
}

var userModel = models.User{
	CountryCode: "+62",
	PhoneNumber: "8123456789",
	Name:        "Test",
	ImageUrl:    "http://test.com",
	Status:      "available",
}

func TestUserService_FindUserById(t *testing.T) {

	userRepository.Mock.On(findUserById, 1).Return(user, nil)

	result, err := userService.FindUserById(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result)
	assert.Equal(t, user.Name, result.Name)
}

func TestUserService_CreateUser(t *testing.T) {
	userRequest := schemas.UserRequest{
		CountryCode: "+62",
		PhoneNumber: "8123456789",
		Name:        "Test",
		ImageUrl:    "http://test.com",
		Status:      "available",
	}

	userRepository.Mock.On(createUser, userModel).Return(user, nil)

	result, err := userService.CreateUser(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.UserId, result.UserId)
}

func TestUserService_UpdateUser(t *testing.T) {
	userRequest := schemas.UpdateUserRequest{
		CountryCode: "+62",
		PhoneNumber: "8123456789",
		Name:        "Test",
		ImageUrl:    "http://test.com",
		Status:      "available",
	}

	userRepository.Mock.On(updateUser, userModel, 1).Return(user, nil)

	result, err := userService.UpdateUser(userRequest, 1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result)
	assert.Equal(t, user.Name, result.Name)
}
