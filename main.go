package main

import (
	"english-ai-go/config"
	"english-ai-go/migrations"
	"log"
)

func main() {
	config.ConnectDatabase()

	migrations.Migrate()

	log.Println("Connected to database!")
}
