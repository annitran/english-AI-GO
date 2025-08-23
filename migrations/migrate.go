package migrations

import (
	"english-ai-go/config"
	"english-ai-go/models"
	"fmt"
	"log"
)

func Migrate() {
	// Tự động tạo bảng
	err := config.GetDB().AutoMigrate(
		&models.User{},
		&models.Chat{},
		&models.Word{},
		&models.History{})

	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
	fmt.Println("Database migrated successfully.")
}
