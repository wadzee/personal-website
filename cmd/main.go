package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"radziramli/handler"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	err := godotenv.Load()

	if err != nil {
		log.Printf("No .env file found, using default value...")
	}

	port := os.Getenv("PORT")

	// server static file on /public path
	app.Static("/public", "static")

	// bind views to the server
	handler.Routes(app)

	// Start server
	go func() {
		if err := app.Start(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
