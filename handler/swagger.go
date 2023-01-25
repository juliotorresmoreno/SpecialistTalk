package handler

import (
	"github.com/labstack/echo/v4"
)

func AttachSwaggerApi(g *echo.Group) *echo.Group {

	g.Static("", "api")

	return g
}
