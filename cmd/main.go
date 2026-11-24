package main

import (
	"errors"
	"log"
	"net/http"

	"project-courier/config"
	"project-courier/internal/handler"
	"project-courier/internal/repository"
	"project-courier/internal/routes"
	"project-courier/internal/usecase"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "project-courier/docs"
)

// @title Courier Management System API - by Kartika Jatilaksono
// @version 1.0
// @description API Documentation for Courier Management System.
// @BasePath /api/v1
func main() {

	// Load Environment Configuration
	config.LoadConfig()

	// Connect Database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()

	log.Println("Database connected successfully")

	// Initialize Repository
	courierRepository := repository.NewCourierRepository(db)
	levelRepository := repository.NewLevelRepository(db)
	statusRepository := repository.NewStatusRepository(db)

	// Initialize Usecase
	courierUsecase := usecase.NewCourierUsecase(
		courierRepository,
		levelRepository,
		statusRepository,
	)

	levelUsecase := usecase.NewLevelUsecase(levelRepository)
	statusUsecase := usecase.NewStatusUsecase(statusRepository)

	// Initialize Handler
	courierHandler := handler.NewCourierHandler(courierUsecase)
	levelHandler := handler.NewLevelHandler(levelUsecase)
	statusHandler := handler.NewStatusHandler(statusUsecase)

	// Initialize Echo
	e := echo.New()

	// Swagger Documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Register Routes
	routes.RegisterRoutes(
		e,
		courierHandler,
		levelHandler,
		statusHandler,
	)

	// Start Server
	port := config.GetAppPort()
	if port == "" {
		log.Fatal("APP_PORT is not configured")
	}

	log.Println("Server running on port:", port)

	if err := e.Start(":" + port); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Failed to start server:", err)
	}
}
