package handlers

import (
	"english-ai-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User not found in context",
		})
		return
	}

	user := userData.(*models.User)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
