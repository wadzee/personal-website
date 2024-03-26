package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/wadzee/personal-website/model"
	"github.com/wadzee/personal-website/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "radzi@radziramli.com",
	}
	return render(c, user.Show(u))
}
