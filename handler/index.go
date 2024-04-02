package handler

import (
	"radziramli/view/pages"

	"github.com/labstack/echo/v4"
)

type IndexPage struct{}

func (h IndexPage) RenderMainPage(c echo.Context) error {

	return render(c, pages.Show())
}
