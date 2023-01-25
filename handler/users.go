package handler

import "github.com/labstack/echo/v4"

type UsersHandler struct {
}

func AttachUsers(g *echo.Group) {
	u := new(UsersHandler)

	g.GET("", u.GET)
}

func (u UsersHandler) GET(c echo.Context) error {
	return c.String(200, "hello world")
}
