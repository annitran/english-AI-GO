package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	CreateHistoryTitle(history *models.History) error
	GetHistoriesByUser(userID uint) ([]models.History, error)
	GetHistoryByID(history *models.History, id string) error
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository() HistoryRepository {
	return &historyRepository{
		db: config.GetDB(),
	}
}

func (r *historyRepository) CreateHistoryTitle(history *models.History) error {
	return r.db.Create(history).Error
}

func (r *historyRepository) GetHistoriesByUser(userID uint) ([]models.History, error) {
	var histories []models.History
	if err := r.db.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *historyRepository) GetHistoryByID(history *models.History, id string) error {
	return r.db.Preload("Chats").First(history, id).Error
}
