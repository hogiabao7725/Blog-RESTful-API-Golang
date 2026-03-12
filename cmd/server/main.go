package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/config"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/database"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/repository"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/routes"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/service"
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

	// 1st floor repository
	userRepo := repository.NewUserRepository(db)

	// 2nd floor service
	userService := service.NewUserService(userRepo)

	// 3rd floor handler
	userHandler := handler.NewUserHandler(userService)
	
	// mux
	mux := http.NewServeMux()

	// routes
	routes.SetupUserRoutes(mux, userHandler)

	// server
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mux,
	}

	fmt.Printf("Server is up and running on PORT %s\n", cfg.ServerPort)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
