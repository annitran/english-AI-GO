package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	CreateHistory(history *models.History) error
	GetHistoriesByUser(userID uint) ([]models.History, error)
	GetHistoryByID(id uint) (*models.History, error)
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository() HistoryRepository {
	return &historyRepository{
		db: config.GetDB(),
	}
}

func (r *historyRepository) CreateHistory(history *models.History) error {
	return r.db.Create(history).Error
}

func (r *historyRepository) GetHistoriesByUser(userID uint) ([]models.History, error) {
	var histories []models.History
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&histories).Error
	return histories, err
}

func (r *historyRepository) GetHistoryByID(id uint) (*models.History, error) {
	var history models.History
	err := r.db.
		Preload("Chats").
		First(&history, id).Error
	return &history, err
}
