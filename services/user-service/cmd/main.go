package main

import (
	"user-service/internal/api"
	"user-service/internal/config"

	"github.com/labstack/echo/v4"
)

var HttpPort = config.GetEnv("HTTP_PORT", "8080")

func main() {
	_ = config.Load()

	// Setup DB Connection

	// Setup gRPC Connection

	// Setup API routes
	e := echo.New()
	api.SetupRoutes(e)

	e.Logger.Fatal(e.Start(HttpPort))
}
