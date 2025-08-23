package handlers

import (
	"english-ai-go/models"
	"english-ai-go/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type historyHandler struct {
	repo repositories.HistoryRepository
}

func NewHistoryHandler(repo repositories.HistoryRepository) *historyHandler {
	return &historyHandler{
		repo: repo,
	}
}

func (h *historyHandler) Create(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var req struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	history := models.History{
		UserID: userID,
		Title:  req.Title,
	}
	if err := h.repo.CreateHistoryTitle(&history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"history": history})
}

// GET /api/v1/histories
func (h *historyHandler) GetList(c *gin.Context) {
	userData, _ := c.Get("user")
	user := userData.(*models.User)

	histories, err := h.repo.GetHistoriesByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get histories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"histories": histories})
}

// GET /api/v1/chat/:id
func (h *historyHandler) GetHistoryByID(c *gin.Context) {
	id := c.Param("id")

	var history models.History
	if err := h.repo.GetHistoryByID(&history, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "History not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"history":  history,
		"messages": history.Chats,
	})
}
