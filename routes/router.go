package routes

import (
	"english-ai-go/handlers"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	messageHandler := handlers.NewMessageHandler()
	registerHandler := handlers.NewRegisterHandler(repositories.NewUserRegister())
	loginHandler := handlers.NewLoginHandler()

	router.GET("/api/v1/message", messageHandler.Get)
	router.POST("/api/v1/message", messageHandler.Create)

	router.POST("/api/v1/register", registerHandler.Create)

	router.POST("/api/v1/login", loginHandler.Login)

	return router
}
