package main

import "english-ai-go/routes"

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
