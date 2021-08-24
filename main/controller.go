package main

import (
	"net/http"

	"github.com/covid19/manager"
	"github.com/covid19/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	repository.SetupMongo()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.PUT("/cases", manager.UpdateCaseCount)
	e.GET("/cases", manager.GetCaseCount)

	e.Logger.Fatal(e.Start(":8080"))
}
