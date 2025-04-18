package routes

import (
	"english-ai-go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	messageHandler := handlers.NewMessageHandler()

	router.GET("/api/v1/message", messageHandler.Get)
	router.POST("/api/v1/message/submit", messageHandler.Creat)

	return router
}
