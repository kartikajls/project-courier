package handler

import "github.com/labstack/echo/v4"

type StatusHandler interface {
	GetAll(c echo.Context) error
	GetByType(c echo.Context) error
}
