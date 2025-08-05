package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"gorm.io/gorm"
)

type ChatRepository interface {
	CreateMessage(chat *models.Chat) error
	GetMessagesByHistoryID(historyID uint) ([]models.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository() ChatRepository {
	return &chatRepository{
		db: config.GetDB(),
	}
}

func (r *chatRepository) CreateMessage(chat *models.Chat) error {
	return r.db.Create(chat).Error
}

func (r *chatRepository) GetMessagesByHistoryID(historyID uint) ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.
		Where("history_id = ?", historyID).
		Order("created_at asc").
		Find(&chats).Error
	return chats, err
}
