package routes

import (
	"english-ai-go/handlers"
	"github.com/gin-gonic/gin"
)

func SetupLoginRouter(r *gin.Engine) {
	loginHandler := handlers.NewLoginHandler()

	r.POST("/api/v1/login", loginHandler.Login)
}
