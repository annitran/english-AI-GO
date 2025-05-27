package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func ConnectDatabase() {
	// Load biến môi trường từ file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Lấy DSN từ biến môi trường
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found in environment variables")
	}

	// Kết nối database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db = database
}
