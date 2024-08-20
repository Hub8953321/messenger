package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	e "messager/src/internal/errors"
	"messager/src/internal/models"
	"net/http"
)

func (h *Handler) CreateChat(ctx echo.Context) error {
	token := ctx.QueryParam("token")

	var data models.ChatCreateDTO

	buf := make([]byte, 1024)

	num, _ := ctx.Request().Body.Read(buf)
	err := json.Unmarshal(buf[:num], &data)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant unmarshal")
	}
	//TODO
	id, err := h.Chat.CreateChat(context.Background(), token, data)

	if err != nil {
		if errors.Is(err, e.UserUnauthorized) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		if errors.Is(err, e.ArrayTooShort) {
			return echo.NewHTTPError(http.StatusBadRequest, "array of members too short")
		}
		if errors.Is(err, e.AccessError) {
			return echo.NewHTTPError(http.StatusBadRequest, "user must be in array of members")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]int{"chatId": id})
}

func (h *Handler) AddMembers(ctx echo.Context) error {
	token := ctx.QueryParam("token")

	var data models.ChatAddMemberDTO

	buf := make([]byte, 1024)

	num, _ := ctx.Request().Body.Read(buf)
	err := json.Unmarshal(buf[:num], &data)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant unmarshal")
	}

	err = h.Chat.AddMembers(context.Background(), token, data)

	if err != nil {
		if errors.Is(err, e.UserUnauthorized) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h* Handler) RemoveMembers(ctx echo.Context) error {
	token := ctx.QueryParam("token")

	var data models.ChatRemoveMemberDTO

	buf := make([]byte, 1024)

	num, _ := ctx.Request().Body.Read(buf)
	err := json.Unmarshal(buf[:num], &data)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant unmarshal")
	}

	err=h.Chat.RemoveMembers(context.Background(), token, data)

	if err != nil {
		if errors.Is(err, e.UserUnauthorized) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		if errors.Is(err, e.AccessError) {
			return echo.NewHTTPError(http.StatusBadRequest, "user must be creator of the chat")
		}
		if errors.Is(err, e.NoneRowsAffected) {
			return echo.NewHTTPError(http.StatusBadRequest, "None of users were removed")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}