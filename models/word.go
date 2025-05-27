package models

import "time"

type Word struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Word      string
	Meaning   string
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID;references:ID"`
}
