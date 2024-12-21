package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tnqbao/gau_validation/config"
	"github.com/tnqbao/gau_validation/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.InitRedis()
	log.Println("Redis initialized successfully")

	db := config.InitDB()
	log.Println("Database initialized successfully")

	router := routes.SetupRouter(db)
	log.Println("Router initialized successfully")

	port := ":8081"
	log.Printf("Server running on port %s", port)
	router.Run(port)
}
