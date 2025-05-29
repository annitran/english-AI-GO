package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type messageHandler struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func NewMessageHandler() *messageHandler {
	return &messageHandler{}
}

func (h *messageHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func (h *messageHandler) Create(c *gin.Context) {
	var req messageHandler
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || req.Age == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Thông tin không hợp lệ!!!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": req.Name + " " + req.Age + " tuổi",
	})
}
