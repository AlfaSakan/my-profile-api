package routes

import (
	"myProfileApi/src/handlers"

	"github.com/gin-gonic/gin"
)

const SESSION_ROUTE = "/session"

func RouteSession(router *gin.RouterGroup, handler *handlers.SessionHandler) {
	router.POST(SESSION_ROUTE, handler.PostSessionHandler)

	router.DELETE(SESSION_ROUTE+"/:sessionId", handler.DeleteSessionHandler)
}
