package repositories

import (
	"fmt"
	"regexp"
	"testing"

	mockModule "github.com/AlfaSakan/my-profile-api.git/src/mock_module"
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository__FindUserById(t *testing.T) {
	db, mock := mockModule.Database()

	newUser := models.User{
		UserId:      "1",
		CountryCode: "+62",
		PhoneNumber: "8123456789",
		Name:        "test name",
		ImageUrl:    "image",
		Status:      "active",
		CreatedAt:   123,
		UpdatedAt:   123,
	}

	rows := sqlmock.NewRows([]string{"user_id", "country_code", "name", "phone_number", "created_at", "updated_at", "image_url", "status"}).
		AddRow(newUser.UserId, newUser.CountryCode, newUser.Name, newUser.PhoneNumber, newUser.CreatedAt, newUser.UpdatedAt, newUser.ImageUrl, newUser.Status)

	t.Run("success find user", func(t *testing.T) {
		userId := "1"
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`user_id` = ?")).
			WithArgs(userId).
			WillReturnRows(rows)

		userRepo := NewUserRepository(db)
		user, errRepo := userRepo.FindUserById(userId)
		assert.Nil(t, errRepo)
		assert.Equal(t, newUser, user)
	})

	t.Run("failed find user by id", func(t *testing.T) {
		userId := 2
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`user_id` = ?")).
			WithArgs(userId).
			WillReturnRows(sqlmock.NewRows(nil))

		userRepo := NewUserRepository(db)
		_, errRepo := userRepo.FindUserById("2")
		assert.Nil(t, errRepo)
	})
}

func TestUserRepository__CreateUser(t *testing.T) {
	db, mock := mockModule.Database()

	newUser := models.User{
		UserId:      "1",
		CountryCode: "+62",
		PhoneNumber: "8123456789",
		Name:        "test name",
		ImageUrl:    "image",
		Status:      "active",
		CreatedAt:   123,
		UpdatedAt:   123,
	}

	t.Run("success create user", func(t *testing.T) {
		userRequest := models.User{
			CountryCode: newUser.CountryCode,
			Name:        newUser.Name,
			PhoneNumber: newUser.PhoneNumber,
			ImageUrl:    newUser.ImageUrl,
			Status:      newUser.Status,
			CreatedAt:   123,
			UpdatedAt:   123,
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`country_code`,`phone_number`,`name`,`image_url`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
			WithArgs(newUser.CountryCode, newUser.PhoneNumber, newUser.Name, newUser.ImageUrl, newUser.Status, newUser.CreatedAt, newUser.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		userRepo := NewUserRepository(db)
		user, errRepo := userRepo.CreateUser(userRequest)
		assert.Nil(t, errRepo)
		assert.Equal(t, newUser.UserId, user.UserId)
		assert.Equal(t, newUser.PhoneNumber, user.PhoneNumber)
	})

	t.Run("failed create user if country code empty", func(t *testing.T) {
		userRequest := models.User{
			CountryCode: "",
			Name:        newUser.Name,
			PhoneNumber: newUser.PhoneNumber,
			ImageUrl:    newUser.ImageUrl,
			Status:      newUser.Status,
			CreatedAt:   123,
			UpdatedAt:   123,
		}

		errMessage := "Country code cannot be empty"

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`country_code`,`phone_number`,`name`,`image_url`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
			WithArgs("", newUser.PhoneNumber, newUser.Name, newUser.ImageUrl, newUser.Status, newUser.CreatedAt, newUser.UpdatedAt).
			WillReturnError(fmt.Errorf(errMessage))
		mock.ExpectRollback()

		userRepo := NewUserRepository(db)
		_, errRepo := userRepo.CreateUser(userRequest)

		errMock := mock.ExpectationsWereMet()
		if errMock != nil {
			assert.Equal(t, errMessage, errMock.Error())
		}

		assert.Error(t, errRepo)

	})
}

func TestUserRepository_UpdateUser(t *testing.T) {
	db, mock := mockModule.Database()

	userRepo := NewUserRepository(db)

	t.Run("success update user", func(t *testing.T) {
		userRequest := models.User{
			Name: "Change name",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users`").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		user, err := userRepo.UpdateUser(userRequest, "1")

		assert.Nil(t, err)
		assert.Equal(t, userRequest.Name, user.Name)
	})
}
