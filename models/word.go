package models

import "time"

type Word struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Word      string    `json:"word"`
	Meaning   string    `json:"meaning"`
	Example   string    `json:"example"`
	IsLearned bool      `json:"is_learned" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}
