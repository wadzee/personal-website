package main

import (
	"fmt"

	"radziramli/handler"

	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()
	fmt.Println("it is working")

	MainPage := handler.IndexPage{}

	app.GET("/", MainPage.RenderMainPage)
	app.Static("/static/css", "static/css")
	app.Logger.Fatal(app.Start(":3000"))
}
