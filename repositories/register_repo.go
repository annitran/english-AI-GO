package repositories

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRegister interface {
	CheckEmail(email string) (*models.User, error)
	CreateUser(name, email, password string) (*models.User, error)
}

type userRegister struct {
	db *gorm.DB
}

func NewUserRegister() UserRegister {
	return &userRegister{
		db: config.GetDB(),
	}
}

func (r *userRegister) CheckEmail(email string) (*models.User, error) {
	var existingUser models.User
	if err := r.db.Where("email = ?", email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // không có user, hợp lệ
		}
		return nil, err
	}
	return &existingUser, nil
}

func (r *userRegister) CreateUser(name, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
