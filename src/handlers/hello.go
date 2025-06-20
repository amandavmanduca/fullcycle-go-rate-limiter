package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
