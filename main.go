package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Tạo router mặc định
	router := gin.Default() // Một *Engine với logger và recovery sẵn

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})

	router.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		c.JSON(http.StatusOK, gin.H{
			"message": name + " " + age + " tuổi",
		})
	})

	// Khởi động http server tại cổng 8080
	router.Run(":8080")
}
