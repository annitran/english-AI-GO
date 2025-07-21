package models

import "time"

type History struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time

	User  User   `gorm:"foreignKey:UserID;references:ID"`
	Chats []Chat `gorm:"foreignKey:HistoryID"`
}
