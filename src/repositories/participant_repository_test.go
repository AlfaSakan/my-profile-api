package repositories

import (
	mockModule "myProfileApi/src/mock_module"
	"myProfileApi/src/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestParticipantRepository_FindAllParticipant(t *testing.T) {
	db, mock := mockModule.Database()

	t.Run("success find list participant", func(t *testing.T) {
		chatRoomId := 1

		participantData := models.Participant{
			UserId:     1,
			ChatRoomId: 1,
		}

		participantData2 := models.Participant{
			UserId:     2,
			ChatRoomId: 1,
		}

		rows := sqlmock.NewRows([]string{"chat_room_id", "user_id"}).
			AddRow(participantData.ChatRoomId, participantData.UserId).
			AddRow(participantData2.ChatRoomId, participantData2.UserId)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `participants` WHERE `participants`.`chat_room_id` = ?")).
			WithArgs(chatRoomId).
			WillReturnRows(rows)

		repo := NewParticipantRepository(db)
		participant, err := repo.FindAllParticipant(chatRoomId)

		assert.Nil(t, err)
		assert.Equal(t, participantData, participant[0])
		assert.Equal(t, participantData2, participant[1])
	})
}

func TestParticipantRepository_FindAllChatRoom(t *testing.T) {
	db, mock := mockModule.Database()

	t.Run("success find list chat room", func(t *testing.T) {
		userId := 1

		participantData := models.Participant{
			UserId:     1,
			ChatRoomId: 1,
		}

		participantData2 := models.Participant{
			UserId:     1,
			ChatRoomId: 2,
		}

		rows := sqlmock.NewRows([]string{"chat_room_id", "user_id"}).
			AddRow(participantData.ChatRoomId, participantData.UserId).
			AddRow(participantData2.ChatRoomId, participantData2.UserId)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `participants` WHERE `participants`.`user_id` = ?")).
			WithArgs(userId).
			WillReturnRows(rows)

		repo := NewParticipantRepository(db)
		participants, err := repo.FindAllChatRoom(uint(userId))

		assert.Nil(t, err)
		assert.Equal(t, participantData, participants[0])
		assert.Equal(t, participantData2, participants[1])
	})
}

func TestParticipantRepository_CreateParticipant(t *testing.T) {
	db, mock := mockModule.Database()

	t.Run("success create new participant", func(t *testing.T) {
		participant := &models.Participant{
			UserId:     1,
			ChatRoomId: 1,
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `participants`").
			WithArgs(participant.ChatRoomId, participant.UserId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := NewParticipantRepository(db)
		err := repo.CreateParticipant(participant)

		assert.Nil(t, err)
	})
}
