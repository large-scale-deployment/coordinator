package main

import (
	"github.com/large-scale-deployment/coordinator/handlers"
	"github.com/large-scale-deployment/coordinator/models"
	"github.com/large-scale-deployment/coordinator/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func startService(ssHandler *handlers.ServiceStatusHandler, psHandler *handlers.PodStatusHandler) {

    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    // Routes
    e.POST("/registry/service_statuses", ssHandler.Create)
    e.POST("/registry/pod_statuses", psHandler.Create)

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}

func main() {
	db, err := gorm.Open(sqlite.Open("coordinator.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	(&models.Models{DB: db}).AutoMigrate()

	reg := &services.Registry{DB: db}
	ssHandler := &handlers.ServiceStatusHandler{Registry: reg}
	psHandler := &handlers.PodStatusHandler{Registry: reg}
	startService(ssHandler, psHandler)
}
