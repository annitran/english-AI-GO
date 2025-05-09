package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type messageHandler struct{}

func NewMessageHandler() *messageHandler {
	return &messageHandler{}
}

func (h *messageHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func (h *messageHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")

	if name == "" || age == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Vui lòng kiểm tra lại thông tin!!!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": name + " " + age + " tuổi",
	})
}
