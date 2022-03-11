package main

import (
	"lss/coordinator/handlers"
	"lss/coordinator/models"
	"lss/coordinator/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func startService(reqHandler *handlers.RequestHandler) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.POST("/registry/statuses", reqHandler.RegisterServiceStatus)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	(&models.Models{DB: db}).AutoMigrate()

	reg := &services.Registry{DB: db}
	handler := &handlers.RequestHandler{Registry: reg}
	startService(handler)
}
