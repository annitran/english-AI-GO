package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserLogin interface {
	AuthenticateUser(email, password string) (*models.User, error)
}

type userLogin struct {
	db *gorm.DB
}

func NewUserLogin() UserLogin {
	return &userLogin{
		db: config.GetDB(),
	}
}

func (r *userLogin) AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}
