package main

import (
	v1 "invoice-management/api/v1"
	"invoice-management/internal/db"
	"invoice-management/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Database
	db.Init()

	// Auto Migrate the database schema
	db.DB.AutoMigrate(&models.Invoice{})

	// Register Routes
	v1.RegisterRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
