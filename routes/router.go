package routes

import (
	"english-ai-go/handlers"
	"english-ai-go/middlewares"
	"english-ai-go/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React app
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	registerHandler := handlers.NewRegisterHandler(repositories.NewUserRegister())
	loginHandler := handlers.NewLoginHandler(repositories.NewUserLogin())
	userHandler := handlers.NewUserHandler(repositories.NewUserRepository())

	messageHandler := handlers.NewMessageHandler(
		repositories.NewChatRepository(),
		repositories.NewHistoryRepository(),
	)
	wordHandler := handlers.NewWordHandler(repositories.NewWordRepository())
	historyHandler := handlers.NewHistoryHandler(repositories.NewHistoryRepository())

	router.POST("/api/v1/register", registerHandler.Create)
	router.POST("/api/v1/login", loginHandler.Login)

	// route xác thực token
	userRepo := repositories.NewUserRepository()
	auth := router.Group("/api/v1", middlewares.AuthToken(userRepo))
	{
		auth.GET("/messages", messageHandler.GetAllByHistoryID)
		auth.POST("/message", messageHandler.Create)

		auth.GET("/user/:id", userHandler.GetUserByID)

		auth.POST("/logout", handlers.Logout)

		auth.POST("/users/words", wordHandler.Create)
		auth.GET("/users/words", wordHandler.GetList)

		auth.GET("/histories", historyHandler.GetList)
		auth.GET("/history/:id", historyHandler.GetByID)
	}

	return router
}
