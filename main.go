package main

import "english-ai-go/routes"

func main() {
	routes.SetupRouter().Run(":8080")
}
