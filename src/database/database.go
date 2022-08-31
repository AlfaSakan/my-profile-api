package database

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"
	"github.com/AlfaSakan/my-profile-api.git/src/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, *services.UserService, *services.MessageService, *services.ChatRoomService, *services.ParticipantService, *services.SessionService, *services.ContactService) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/myProfile?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ChatRoom{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Participant{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Contact{})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	messageRepository := repositories.NewMessageRepository(db)
	messageService := services.NewMessageService(messageRepository)

	participantRepository := repositories.NewParticipantRepository(db)
	participantService := services.NewParticipantService(participantRepository, userRepository)

	chatRoomRepository := repositories.NewChatRoomRepository(db)
	chatRoomService := services.NewChatRoomService(chatRoomRepository, participantRepository, messageRepository)

	sessionRepository := repositories.NewSessionRepository(db)
	sessionService := services.NewSessionService(sessionRepository, userRepository)

	contactRepository := repositories.NewContactRepository(db)
	contactService := services.NewContactService(contactRepository)

	return db, userService, messageService, chatRoomService, participantService, sessionService, contactService
}
