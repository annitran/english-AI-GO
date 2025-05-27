package routes

import (
	"english-ai-go/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRegisterRouter(r *gin.Engine) {
	registerHandler := handlers.NewRegisterHandler()

	r.POST("/api/v1/register", registerHandler.Create)
}
