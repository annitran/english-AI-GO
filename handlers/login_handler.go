package handlers

import (
	"english-ai-go/middlewares"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginHandler struct {
	repo repositories.UserLogin
}

func NewLoginHandler(repo repositories.UserLogin) *loginHandler {
	return &loginHandler{
		repo: repo,
	}
}

func (h *loginHandler) Login(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid information!",
		})
		return
	}

	user, err := h.repo.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Incorrect email or password!",
		})
		return
	}

	token, err := middlewares.GenerateToken(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot login!!!",
		})
		return
	}

	c.SetCookie(
		"token",
		token,
		3600/60*5,
		"/",
		"",
		false, // secure = false vì đang dev (true nếu chạy https)
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!!!",
		"user":    user,
		"token":   token,
	})
}
