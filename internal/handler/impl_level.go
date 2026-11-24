package handler

import (
	"errors"
	"net/http"

	"project-courier/internal/usecase"

	"github.com/labstack/echo/v4"
)

type levelHandler struct {
	levelUsecase usecase.LevelUsecase
}

func NewLevelHandler(levelUsecase usecase.LevelUsecase) LevelHandler {
	return &levelHandler{
		levelUsecase: levelUsecase,
	}
}

// Get All Levels
// @Summary Get all levels
// @Description Get all courier levels
// @Tags Levels
// @Produce json
// @Success 200 {array} dto.LevelResponse
// @Failure 500 {object} map[string]interface{}
// @Router /levels [get]
func (h *levelHandler) GetAll(c echo.Context) error {
	result, err := h.levelUsecase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve levels",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Levels retrieved successfully",
		"data":    result,
	})
}

func (h *levelHandler) GetByID(c echo.Context) error {
	levelID := c.Param("level_id")

	result, err := h.levelUsecase.GetByID(c.Request().Context(), levelID)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		if errors.Is(err, usecase.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Level not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve level",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Level retrieved successfully",
		"data":    result,
	})
}
