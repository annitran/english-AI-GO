package handlers

import (
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
	"net/http"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerHandler struct {
	repo repositories.UserRegister
}

func NewRegisterHandler(repo repositories.UserRegister) *registerHandler {
	return &registerHandler{
		repo: repo,
	}
}

func (h *registerHandler) Create(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid information!!!",
		})
		return
	}

	user, err := h.repo.CheckEmail(req.Email)
	if err != nil {
		// lỗi DB 500
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if user != nil {
		// email tồn tại
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email already exists!",
		})
		return
	}

	user, err = h.repo.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot register!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered successful!!!",
		"user":    user,
	})
}
