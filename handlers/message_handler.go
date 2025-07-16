package handlers

import (
	"english-ai-go/AiServices"
	"english-ai-go/models"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type createMessageRequest struct {
	Message   string `json:"message" binding:"required"`
	HistoryID uint   `json:"history_id"`
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
	historyID := req.HistoryID

	// Nếu không có history thì tạo mới
	if historyID == 0 {
		newHistory := models.History{
			Title:  req.Message, // Dùng message đầu tiên làm title
			UserID: user_id,
		}
		if err := h.repo.CreateHistory(&newHistory); err != nil {
			log.Println("Error creating history:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create history",
			})
			return
		}
		historyID = newHistory.ID
	}

	// Lưu tin nhắn người dùng
	userMsg := models.Chat{
		UserID:    user_id,
		Message:   req.Message,
		IsBot:     false,
		HistoryID: historyID,
	}
	if err := h.repo.CreateMessage(&userMsg); err != nil {
		log.Println("Error saving user message:", err)
	}

	// Gọi AI bot để phản hồi
	botReply, err := AiServices.Reply(req.Message)
	if err != nil {
		log.Println("AI bot is not responding:", err)
		botReply = "Sorry, I couldn't respond right now."
	}

	botMsg := models.Chat{
		UserID:    user_id,
		Message:   botReply,
		IsBot:     true,
		HistoryID: historyID,
	}
	if err := h.repo.CreateMessage(&botMsg); err != nil {
		log.Println("Error saving bot reply:", err)
	}

	messages, _ := h.repo.GetMessagesByHistoryID(historyID)
	c.JSON(http.StatusOK, gin.H{
		"messages":   messages,
		"history_id": historyID,
	})
}

func (h *messageHandler) GetAll(c *gin.Context) {
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
