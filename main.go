package main

import (
	"myProfileApi/src/database"
	"myProfileApi/src/handlers"
	"myProfileApi/src/middlewares"
	"myProfileApi/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	viper.SetConfigFile(".env")

	v1 := router.Group("/v1")
	v1.Use(middlewares.CORSMiddleware())

	db, userService, messageService, chatRoomService, participantService, sessionService := database.DatabaseConnection()
	v1.Use(middlewares.DeserializeUser(db))

	userHandlers := handlers.NewUserHandler(userService)
	routes.RouteUser(v1, userHandlers)

	messageHandlers := handlers.NewMessageHandler(messageService)
	routes.RouteMessage(v1, messageHandlers)

	chatRoomHandlers := handlers.NewChatRoomHandler(chatRoomService, participantService)
	routes.ChatRoomRoute(v1, chatRoomHandlers)

	sessionHandlers := handlers.NewSessionHandler(sessionService, userService)
	routes.RouteSession(v1, sessionHandlers)

	router.Run(":8081")
}
