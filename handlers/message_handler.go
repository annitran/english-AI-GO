package handlers

import (
	"english-ai-go/AiServices"
	"english-ai-go/models"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type createMessageRequest struct {
	Message   string `json:"message" binding:"required"`
	HistoryID uint   `json:"history_id"`
}

type messageHandler struct {
	chatRepo    repositories.ChatRepository
	historyRepo repositories.HistoryRepository
}

func NewMessageHandler(chatRepo repositories.ChatRepository, historyRepo repositories.HistoryRepository) *messageHandler {
	return &messageHandler{
		chatRepo:    chatRepo,
		historyRepo: historyRepo,
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

	// Nếu không có history thì tạo mới
	historyID := req.HistoryID
	if historyID == 0 {
		newHistory := models.History{
			Title:  req.Message, // Dùng message đầu tiên làm title
			UserID: user_id,
		}
		if err := h.historyRepo.CreateHistoryTitle(&newHistory); err != nil {
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
	if err := h.chatRepo.CreateMessage(&userMsg); err != nil {
		log.Println("Error saving user message:", err)
	}

	// Gọi AI bot để phản hồi
	ai := AiServices.NewAIService()
	botReply, err := ai.Reply(req.Message)
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
	if err := h.chatRepo.CreateMessage(&botMsg); err != nil {
		log.Println("Error saving bot reply:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"messages":   []models.Chat{userMsg, botMsg},
		"history_id": historyID,
	})
}

func (h *messageHandler) GetAllByHistoryID(c *gin.Context) {
	historyID_str := c.Query("history_id")
	historyID, err := strconv.Atoi(historyID_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid HistoryID"})
		return
	}

	messages, err := h.chatRepo.GetMessagesByHistoryID(uint(historyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get messages"},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
