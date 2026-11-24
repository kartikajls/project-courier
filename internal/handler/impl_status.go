package handler

import (
	"errors"
	"net/http"

	"project-courier/internal/usecase"

	"github.com/labstack/echo/v4"
)

type statusHandler struct {
	statusUsecase usecase.StatusUsecase
}

func NewStatusHandler(statusUsecase usecase.StatusUsecase) StatusHandler {
	return &statusHandler{
		statusUsecase: statusUsecase,
	}
}

// Get All Statuses
// @Summary Get all statuses
// @Description Get all available courier statuses
// @Tags Statuses
// @Produce json
// @Success 200 {array} dto.StatusResponse
// @Failure 500 {object} map[string]interface{}
// @Router /statuses [get]
func (h *statusHandler) GetAll(c echo.Context) error {
	result, err := h.statusUsecase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve statuses",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Statuses retrieved successfully",
		"data":    result,
	})
}

func (h *statusHandler) GetByType(c echo.Context) error {
	statusType := c.Param("type")

	result, err := h.statusUsecase.GetByType(c.Request().Context(), statusType)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		if errors.Is(err, usecase.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Status not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve status",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Status retrieved successfully",
		"data":    result,
	})
}
