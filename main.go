package main

import (
	"net/http"

	"github.com/juliotorresmoreno/freelive/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	api := e.Group("/api/v1")
	handlers.AttachUsers(api.Group("/users"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
