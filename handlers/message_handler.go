package handlers

import (
	"english-ai-go/models"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type createMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

type messageHandler struct {
	repo repositories.ChatRepository
}

func NewMessageHandler(repo repositories.ChatRepository) *messageHandler {
	return &messageHandler{
		repo: repo,
	}
}

func (h *messageHandler) Create(c *gin.Context) {
	var req createMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid message!",
		})
		return
	}

	user, _ := c.Get("user") // middleware đã auth trước đó
	user_id := user.(*models.User).ID

	// Lưu tin nhắn người dùng
	chat := models.Chat{
		UserID:  user_id,
		Message: req.Message,
		IsBot:   false,
	}
	if err := h.repo.CreateMessage(&chat); err != nil {
		log.Println("Error saving user message:", err)
	}

	// Phản hồi bot giả lập
	reply := models.Chat{
		UserID:  user_id,
		Message: "AI bot replied !!!",
		IsBot:   true,
	}
	if err := h.repo.CreateMessage(&reply); err != nil {
		log.Println("Error saving bot reply:", err)
	}

	messages, _ := h.repo.GetMessagesByUser(user_id)
	c.JSON(http.StatusOK, messages)
}

func (h *messageHandler) Get(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	user, ok := u.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast user"})
		return
	}

	user_id := user.ID

	messages, err := h.repo.GetMessagesByUser(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch messages"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
