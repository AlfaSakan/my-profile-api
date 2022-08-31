package main

import (
	"github.com/AlfaSakan/my-profile-api.git/src/database"
	"github.com/AlfaSakan/my-profile-api.git/src/handlers"
	"github.com/AlfaSakan/my-profile-api.git/src/middlewares"
	"github.com/AlfaSakan/my-profile-api.git/src/routes"
	"github.com/AlfaSakan/my-profile-api.git/src/websocket"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	viper.SetConfigFile(".env")

	router.SetTrustedProxies([]string{"http://localhost:3000"})
	router.Use(cors.AllowAll())

	v1 := router.Group("/v1")
	hub := websocket.NewHub()
	go hub.Run()

	db, userService, messageService, chatRoomService, participantService, sessionService, contactService := database.DatabaseConnection()
	v1.Use(middlewares.DeserializeUser(db))
	// v1.Use(middlewares.CORSMiddleware(), middlewares.DeserializeUser(db))

	userHandlers := handlers.NewUserHandler(userService)
	routes.RouteUser(v1, userHandlers)

	messageHandlers := handlers.NewMessageHandler(messageService, chatRoomService)
	routes.RouteMessage(v1, messageHandlers)

	chatRoomHandlers := handlers.NewChatRoomHandler(chatRoomService, participantService)
	routes.ChatRoomRoute(v1, chatRoomHandlers, hub)

	sessionHandlers := handlers.NewSessionHandler(sessionService, userService)
	routes.RouteSession(v1, sessionHandlers)

	routes.ParticipantRoute(v1, chatRoomHandlers)

	contactHandler := handlers.NewContactHandler(contactService)
	routes.ContactRoute(v1, contactHandler)

	router.Run(":8081")
}
