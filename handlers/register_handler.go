package handlers

import (
	"english-ai-go/config"
	"english-ai-go/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type registerHandler struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewRegisterHandler() *registerHandler {
	return &registerHandler{}
}

func (h *registerHandler) Create(c *gin.Context) {
	var req registerHandler
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Thông tin không hợp lệ!!!",
		})
		return
	}

	db := config.GetDB()

	// Kiểm tra trùng email hay ko?
	var existingUser models.User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email đã tồn tại!",
		})
		return
	}

	// Mã hoá mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Không thể tạo tài khoản!!!",
		})
		return
	}

	// Tạo User mới
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword), // Lưu chuỗi đã mã hoá
	}

	// Lưu vào DB
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Không thể lưu tài khoản!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tạo tài khoản thành công!!!",
		"user":    user,
	})
}
