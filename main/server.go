package main

import (
	"net/http"

	"github.com/covid19/controllers"
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
	e.PUT("/cases", controllers.UpdateCaseCount)
	e.GET("/cases", controllers.GetCaseCount)

	e.Logger.Fatal(e.Start(":8080"))
}
