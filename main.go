package main

import (
	"myProfileApi/src/database"
	"myProfileApi/src/handlers"
	"myProfileApi/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userService, messageService := database.DatabaseConnection()

	userHandlers := handlers.NewUserHandler(userService)
	routes.RouteUser(router, userHandlers)

	messageHandlers := handlers.NewMessageHandler(messageService)
	routes.RouteMessage(router, messageHandlers)

	router.Run(":8081")
}
