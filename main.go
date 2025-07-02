package main

import (
	"english-ai-go/config"
	"english-ai-go/migrations"
	"english-ai-go/routes"
	"log"
)

func main() {
	config.ConnectDatabase()

	migrations.Migrate()

	log.Println("Connected to database!")

	routes.SetupRouter().Run(":8080")
}
