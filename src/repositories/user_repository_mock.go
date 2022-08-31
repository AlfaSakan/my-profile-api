package repositories

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"github.com/stretchr/testify/mock"
)

type IUserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *IUserRepositoryMock) FindUserById(userId string) (models.User, error) {
	arguments := repository.Mock.Called(userId)

	if arguments.Get(0) == nil {
		return models.User{}, nil
	}

	user := arguments.Get(0).(models.User)
	return user, nil
}

func (repository *IUserRepositoryMock) CreateUser(user models.User) (models.User, error) {
	arguments := repository.Mock.Called(user)

	if arguments.Get(0) == nil {
		return user, nil
	}

	user = arguments.Get(0).(models.User)
	return user, nil
}

func (repository *IUserRepositoryMock) UpdateUser(user models.User, userId string) (models.User, error) {
	arguments := repository.Mock.Called(user, userId)

	if arguments.Get(0) == nil {
		return user, nil
	}

	user = arguments.Get(0).(models.User)
	return user, nil
}

func (repository *IUserRepositoryMock) FindUser(user *models.User) (*models.User, error) {
	foundUser := &models.User{}

	return foundUser, nil
}

func (repository *IUserRepositoryMock) SearchNameUser(name string) (*[]models.User, error) {
	users := &[]models.User{}

	return users, nil
}
