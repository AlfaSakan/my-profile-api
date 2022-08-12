package main

import (
	"myProfileApi/src/database"
	"myProfileApi/src/handlers"
	"myProfileApi/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.Use(handlers.CORSMiddleware())

	userService, messageService, chatRoomService, participantService := database.DatabaseConnection()

	userHandlers := handlers.NewUserHandler(userService)
	routes.RouteUser(v1, userHandlers)

	messageHandlers := handlers.NewMessageHandler(messageService)
	routes.RouteMessage(v1, messageHandlers)

	chatRoomHandlers := handlers.NewChatRoomHandler(chatRoomService, participantService)
	routes.ChatRoomRoute(v1, chatRoomHandlers)

	router.Run(":8081")
}
