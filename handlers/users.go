package handlers

import "github.com/labstack/echo/v4"

type UsersController struct {
}

func AttachUsers(g *echo.Group) {
	u := &UsersController{}

	g.GET("", u.Get)
}

func (u UsersController) Get(c echo.Context) error {
	return c.String(200, "hello")
}
