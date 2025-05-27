package routes

import (
	"english-ai-go/handlers"
	"github.com/gin-gonic/gin"
)

func SetupMessageRouter(r *gin.Engine) {
	messageHandler := handlers.NewMessageHandler()

	r.GET("/api/v1/message", messageHandler.Get)
	r.POST("/api/v1/message", messageHandler.Create)
}
