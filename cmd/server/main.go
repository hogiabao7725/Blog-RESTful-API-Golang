package main

import (
	"log"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/config"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/database"
)

func main() {

	// config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// database
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

}
