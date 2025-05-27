package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	SetupMessageRouter(router)

	SetupRegisterRouter(router)

	SetupLoginRouter(router)

	return router
}
