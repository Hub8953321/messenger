package handler

import "github.com/labstack/echo/v4"

func (h *Handler) SingIn(c echo.Context) error {
	h.logger.Info("Hello, World!")
	return c.String(200, "Hello, World!")
}
