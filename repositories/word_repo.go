package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"gorm.io/gorm"
)

type WordRepository interface {
	CreateWord(word *models.Word) error
	IsWordExist(userID uint, word string) (bool, error)
	GetWordList(userID uint) ([]models.Word, error)
}

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository() WordRepository {
	return &wordRepository{
		db: config.GetDB(),
	}
}

func (r *wordRepository) CreateWord(word *models.Word) error {
	return r.db.Create(word).Error
}

func (r *wordRepository) IsWordExist(userID uint, word string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Word{}).
		Where("user_id = ? AND word = ?", userID, word).
		Count(&count).Error
	return count > 0, err
}

func (r *wordRepository) GetWordList(userID uint) ([]models.Word, error) {
	var words []models.Word
	err := r.db.Where("user_id = ?", userID).Find(&words).Error
	return words, err
}
