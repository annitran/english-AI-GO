package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time

	Chats     []Chat    `gorm:"foreignKey:UserID"`
	Words     []Word    `gorm:"foreignKey:UserID"`
	Histories []History `gorm:"foreignKey:UserID"`
}
