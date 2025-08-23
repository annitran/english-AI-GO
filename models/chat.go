package models

import "time"

type Chat struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	HistoryID uint      `gorm:"history_id"`
	Message   string    `json:"message"`
	IsBot     bool      `json:"is_bot"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User    User    `gorm:"foreignKey:UserID;references:ID"`
	History History `gorm:"foreignKey:HistoryID;references:ID"`
}
