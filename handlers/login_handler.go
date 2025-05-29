package handlers

import (
	"english-ai-go/config"
	"english-ai-go/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginHandler struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginHandler() *loginHandler {
	return &loginHandler{}
}

func (h *loginHandler) Login(c *gin.Context) {
	var req loginHandler
	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Thông tin không hợp lệ!",
		})
		return
	}

	db := config.GetDB()

	user := models.User{}

	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email hoặc Mật khẩu không đúng!",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email hoặc Mật khẩu không đúng!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Đăng nhập thành công!",
		"user":    user,
	})
}
