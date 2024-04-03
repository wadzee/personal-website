package handler

import (
	"radziramli/view/pages"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		pages.HomePage().Render(c.Request().Context(), c.Response())
		return nil
	})
}
