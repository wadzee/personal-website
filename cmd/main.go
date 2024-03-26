package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/wadzee/personal-website/handler"
)

type DB struct{}

func main() {
	app := echo.New()
	fmt.Println("it is working")

	userHandler := handler.UserHandler{}

	app.GET("/", userHandler.HandleUserShow)

	app.Logger.Fatal(app.Start(":3000"))
}
