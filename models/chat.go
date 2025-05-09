package models

import "time"

type Chat struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Message   string
	IsBot     bool
	CreatedAt time.Time
	UpdatedAt time.Time

	User    User    `gorm:"foreignKey:UserID;references:ID"`
	History History `gorm:"foreignKey:HistoryID;references:ID"`
}
