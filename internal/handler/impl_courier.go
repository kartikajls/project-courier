package handler

import (
	"errors"
	"net/http"
	"strconv"

	"project-courier/internal/dto"
	"project-courier/internal/usecase"

	"github.com/labstack/echo/v4"
)

type courierHandler struct {
	courierUsecase usecase.CourierUsecase
}

func NewCourierHandler(
	courierUsecase usecase.CourierUsecase,
) CourierHandler {
	return &courierHandler{
		courierUsecase: courierUsecase,
	}
}

// Create Courier
// @Summary Create a new courier
// @Description Create a new courier record
// @Tags Couriers
// @Accept json
// @Produce json
// @Param request body dto.CreateCourierRequest true "Courier data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /couriers [post]
func (h *courierHandler) Create(c echo.Context) error {
	var request dto.CreateCourierRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	result, err := h.courierUsecase.Create(c.Request().Context(), request)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create courier",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Courier created successfully",
		"data":    result,
	})
}

// Get All Couriers
// @Summary Get all couriers
// @Description Get couriers with search, filter, sorting, and pagination
// @Tags Couriers
// @Produce json
// @Param search query string false "Search courier name"
// @Param level query string false "Filter by level, example: 2,3"
// @Param status query string false "Filter by status: active or inactive"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Param sort query string false "Sort field: name or created_at"
// @Param order query string false "Sort order: asc or desc"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /couriers [get]
func (h *courierHandler) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("per_page"))

	request := dto.CourierQueryRequest{
		Search:  c.QueryParam("search"),
		Level:   c.QueryParam("level"),
		Status:  c.QueryParam("status"),
		Page:    page,
		PerPage: perPage,
		Sort:    c.QueryParam("sort"),
		Order:   c.QueryParam("order"),
	}

	result, err := h.courierUsecase.GetAll(c.Request().Context(), request)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve couriers",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Couriers retrieved successfully",
		"data":       result.Data,
		"pagination": result.Pagination,
	})
}

// Get Courier By ID
// @Summary Get courier by ID
// @Description Get detail of a courier
// @Tags Couriers
// @Produce json
// @Param user_id path string true "Courier ID"
// @Success 200 {object} dto.CourierResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /couriers/{user_id} [get]
func (h *courierHandler) GetByID(c echo.Context) error {
	userID := c.Param("user_id")

	result, err := h.courierUsecase.GetByID(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		if errors.Is(err, usecase.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Courier not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve courier",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Courier retrieved successfully",
		"data":    result,
	})
}

// Update Courier
// @Summary Update courier
// @Description Update an existing courier
// @Tags Couriers
// @Accept json
// @Produce json
// @Param user_id path string true "Courier ID"
// @Param request body dto.UpdateCourierRequest true "Updated courier data"
// @Success 200 {object} dto.CourierResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /couriers/{user_id} [put]
func (h *courierHandler) Update(c echo.Context) error {
	userID := c.Param("user_id")

	var request dto.UpdateCourierRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	result, err := h.courierUsecase.Update(c.Request().Context(), userID, request)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		if errors.Is(err, usecase.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Courier not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update courier",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Courier updated successfully",
		"data":    result,
	})
}

// Delete Courier
// @Summary Delete courier
// @Description Delete a courier by ID
// @Tags Couriers
// @Produce json
// @Param user_id path string true "Courier ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /couriers/{user_id} [delete]
func (h *courierHandler) Delete(c echo.Context) error {
	userID := c.Param("user_id")

	err := h.courierUsecase.Delete(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		if errors.Is(err, usecase.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Courier not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete courier",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Courier deleted successfully",
	})
}
