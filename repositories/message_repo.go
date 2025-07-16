package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"gorm.io/gorm"
)

type ChatRepository interface {
	CreateMessage(chat *models.Chat) error
	GetMessagesByUser(userID uint) ([]models.Chat, error)
	GetMessagesByHistoryID(historyID uint) ([]models.Chat, error)
	CreateHistory(history *models.History) error
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

func (r *chatRepository) GetMessagesByUser(userID uint) ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at asc").
		Find(&chats).Error
	return chats, err
}

func (r *chatRepository) GetMessagesByHistoryID(historyID uint) ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.
		Where("history_id = ?", historyID).
		Order("created_at asc").
		Find(&chats).Error
	return chats, err
}

func (r *chatRepository) CreateHistory(history *models.History) error {
	return r.db.Create(history).Error
}
