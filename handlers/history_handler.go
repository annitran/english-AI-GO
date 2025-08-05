package handlers

import (
	"english-ai-go/models"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type createHistoryRequest struct {
	Title string `json:"title binding:required"`
}

type historyHandler struct {
	repo repositories.HistoryRepository
}

func NewHistoryHandler(repo repositories.HistoryRepository) *historyHandler {
	return &historyHandler{
		repo: repo,
	}
}

func (h *historyHandler) Create(c *gin.Context) {
	var req createHistoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid title"})
		return
	}

	user, _ := c.Get("user")
	userID := user.(*models.User).ID

	history := models.History{
		UserID: userID,
		Title:  req.Title,
	}

	if err := h.repo.CreateHistoryTitle(&history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create history",
		})
		return
	}

	c.JSON(http.StatusCreated, history)
}

func (h *historyHandler) GetList(c *gin.Context) {
	user, _ := c.Get("user")
	userID := user.(*models.User).ID

	histories, err := h.repo.GetHistoriesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get history list",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"histories": histories})
}

func (h *historyHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid history ID"})
		return
	}

	history, err := h.repo.GetHistoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"history": history,
	})
}
