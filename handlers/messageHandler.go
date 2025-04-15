package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct{}

func NewMessageHandler() handler {
	return handler{}
}

func (h *handler) CreatGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func (h *handler) CreatPost(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")

	c.JSON(http.StatusOK, gin.H{
		"message": name + " " + age + " tuổi",
	})
}
