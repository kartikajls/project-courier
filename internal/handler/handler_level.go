package handler

import "github.com/labstack/echo/v4"

type LevelHandler interface {
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
}
