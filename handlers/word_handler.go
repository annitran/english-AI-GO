package handlers

import (
	"english-ai-go/models"
	"english-ai-go/repositories"

	"github.com/gin-gonic/gin"
	"net/http"
)

type createWordRequest struct {
	Word    string `json:"word" binding:"required"`
	Meaning string `json:"meaning"`
	Example string `json:"example"`
}

type wordHandler struct {
	repo repositories.WordRepository
}

func NewWordHandler(repo repositories.WordRepository) *wordHandler {
	return &wordHandler{
		repo: repo,
	}
}

func (h *wordHandler) Create(c *gin.Context) {
	var req createWordRequest

	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found!",
		})
		return
	}

	user := userData.(*models.User)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid information",
		})
		return
	}

	word := models.Word{
		UserID:  user.ID,
		Word:    req.Word,
		Meaning: req.Meaning,
		Example: req.Example,
	}

	// kiểm tra word đã tồn tại chưa?
	if exists, err := h.repo.IsWordExist(user.ID, req.Word); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error checking word!",
		})
		return
	} else if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Word already exists!",
		})
		return
	}

	if err := h.repo.CreateWord(&word); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create word!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Word created successfully!",
		"word":    word,
	})
}

func (h *wordHandler) GetList(c *gin.Context) {
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found!",
		})
		return
	}
	user := userData.(*models.User)

	words, err := h.repo.GetWordList(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get words!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"words": words,
	})
}
