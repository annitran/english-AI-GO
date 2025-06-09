package routes

import (
	"english-ai-go/handlers"
	"english-ai-go/middlewares"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	messageHandler := handlers.NewMessageHandler()
	registerHandler := handlers.NewRegisterHandler(repositories.NewUserRegister())
	loginHandler := handlers.NewLoginHandler(repositories.NewUserLogin())
	userRepo := repositories.NewUserRepository()

	router.POST("/api/v1/register", registerHandler.Create)

	router.POST("/api/v1/login", loginHandler.Login)

	auth := router.Group("/api/v1")
	auth.Use(middlewares.AuthToken(userRepo))
	{
		auth.GET("/message", messageHandler.Get)
		auth.POST("/message", messageHandler.Create)

		auth.GET("/user", handlers.GetUser)
	}

	return router
}
