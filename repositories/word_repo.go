package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"gorm.io/gorm"
)

type WordRepository interface {
	CreateWord(word *models.Word) error
	IsWordExist(userID uint, word string) (bool, error)
	GetWords(userID uint, search string) ([]models.Word, error)
	UpdateWord(id uint, data map[string]interface{}) (*models.Word, error)
	Delete(id uint, userID uint) error
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

func (r *wordRepository) GetWords(userID uint, search string) ([]models.Word, error) {
	var words []models.Word
	query := r.db.Where("user_id = ?", userID)
	if search != "" {
		query = query.Where("word LIKE ?", "%"+search+"%")
	}
	if err := query.Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}

func (r *wordRepository) UpdateWord(id uint, data map[string]interface{}) (*models.Word, error) {
	var word models.Word
	if err := r.db.First(&word, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&word).Updates(data).Error; err != nil {
		return nil, err
	}

	return &word, nil
}

func (r *wordRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Word{}).Error
}
