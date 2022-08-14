package database

import (
	"myProfileApi/src/models"
	"myProfileApi/src/repositories"
	"myProfileApi/src/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, *services.UserService, *services.MessageService, *services.ChatRoomService, *services.ParticipantService, *services.SessionService) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/myProfile?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ChatRoom{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.MessageRead{})
	db.AutoMigrate(&models.Participant{})
	db.AutoMigrate(&models.Session{})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	messageRepository := repositories.NewMessageRepository(db)
	messageService := services.NewMessageService(messageRepository)

	participantRepository := repositories.NewParticipantRepository(db)
	participantService := services.NewParticipantService(participantRepository)

	chatRoomRepository := repositories.NewChatRoomRepository(db)
	chatRoomService := services.NewChatRoomService(chatRoomRepository, participantRepository)

	sessionRepository := repositories.NewSessionRepository(db)
	sessionService := services.NewSessionService(sessionRepository, userRepository)

	return db, userService, messageService, chatRoomService, participantService, sessionService
}
