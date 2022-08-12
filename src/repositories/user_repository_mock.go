package repositories

import (
	"myProfileApi/src/models"

	"github.com/stretchr/testify/mock"
)

type IUserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *IUserRepositoryMock) FindUserById(userId int) (models.User, error) {
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

func (repository *IUserRepositoryMock) UpdateUser(user models.User, userId int) (models.User, error) {
	arguments := repository.Mock.Called(user, userId)

	if arguments.Get(0) == nil {
		return user, nil
	}

	user = arguments.Get(0).(models.User)
	return user, nil
}
