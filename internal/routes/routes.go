package routes

import (
	"project-courier/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	courierHandler handler.CourierHandler,
	levelHandler handler.LevelHandler,
	statusHandler handler.StatusHandler,
) {
	// API Version 1
	api := e.Group("/api/v1")

	// =========================
	// COURIER ROUTES
	// =========================
	couriers := api.Group("/couriers")

	couriers.POST("", courierHandler.Create)
	couriers.GET("", courierHandler.GetAll)
	couriers.GET("/:user_id", courierHandler.GetByID)
	couriers.PUT("/:user_id", courierHandler.Update)
	couriers.DELETE("/:user_id", courierHandler.Delete)

	// =========================
	// LEVEL ROUTES (READ ONLY)
	// =========================
	levels := api.Group("/levels")

	levels.GET("", levelHandler.GetAll)
	// levels.GET("/:level_id", levelHandler.GetByID)

	// =========================
	// STATUS ROUTES (READ ONLY)
	// =========================
	statuses := api.Group("/statuses")

	statuses.GET("", statusHandler.GetAll)
	// statuses.GET("/:type", statusHandler.GetByType)
}
