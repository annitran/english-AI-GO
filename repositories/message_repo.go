package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"gorm.io/gorm"
)

type ChatRepository interface {
	CreateMessage(message *models.Chat) error
	GetMessagesByHistory(historyID uint) ([]models.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository() ChatRepository {
	return &chatRepository{
		db: config.GetDB(),
	}
}

func (r *chatRepository) CreateMessage(message *models.Chat) error {
	return r.db.Create(message).Error
}

func (r *chatRepository) GetMessagesByHistory(historyID uint) ([]models.Chat, error) {
	var messages []models.Chat
	err := r.db.
		Where("history_id = ?", historyID).
		Order("created_at asc").
		Find(&messages).Error
	return messages, err
}
