package main

import (
	"context"
	"log"
	"user-service/internal/api"
	"user-service/internal/config"
	"user-service/internal/db"

	"github.com/labstack/echo/v4"
)

var HttpPort = config.GetEnv("HTTP_PORT", "8001")

func main() {
	// Setup Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load Config
	cfg := config.Load()

	// Setup DB Connection
	_, err := db.NewPostgresDB(ctx, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to users db")

	// Setup gRPC Connection

	// Setup API routes
	e := echo.New()
	api.SetupRoutes(e)
	
	e.Logger.Fatal(e.Start(HttpPort))
}
