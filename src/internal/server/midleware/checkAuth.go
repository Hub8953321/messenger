package midleware

import (
	"github.com/labstack/echo/v4"
	"messager/src/internal/handler"
	"net/http"
)

func CheckAuth(next echo.HandlerFunc, handler *handler.Handler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		access := ctx.Request().Header.Get("Authorization")
		if len(access) == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		id, err := handler.AccessConverter.ParseToken(access)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		ctx.Set("id", id)

		return next(ctx)
	}
}
