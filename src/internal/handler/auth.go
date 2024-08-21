package handler

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	e "messager/src/internal/errors"
	"messager/src/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) SingUp(ctx echo.Context) error {
	var dto models.UserSingUpDTO

	buf := make([]byte, 1024)

	num, _ := ctx.Request().Body.Read(buf)
	err := json.Unmarshal(buf[:num], &dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, access, refresh, err := h.Auth.SignUp(ctx.Request().Context(), dto)
	if err != nil {
		if errors.Is(err, e.LoginBusy) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{"id": id, "access": access, "refresh": refresh})
}

func (h *Handler) SignIn(ctx echo.Context) error {
	var dto models.UserSignInDTO
	buf := make([]byte, 1024)

	num, _ := ctx.Request().Body.Read(buf)
	err := json.Unmarshal(buf[:num], &dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, access, refresh, err := h.Auth.SignIn(ctx.Request().Context(), dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"id": id, "access": access, "refresh": refresh})
}

func (h *Handler) Refresh(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token := ctx.QueryParam("token")

	access, refresh, err := h.Auth.Refresh(ctx.Request().Context(), id, token)
	if err != nil {
		if errors.Is(err, e.UserIsAbsent) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"access": access, "refresh": refresh})
}
