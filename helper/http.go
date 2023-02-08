package helper

import (
	"strconv"

	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/labstack/echo/v4"
)

func Paginate(c echo.Context) (int, int) {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	skip, _ := strconv.Atoi(c.QueryParam("skip"))
	if limit == 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	return limit, skip
}

func ValidateSession(c echo.Context) (*model.User, error) {
	session := c.Get("session")
	if session == nil {
		return nil, echo.NewHTTPError(401, "unauthorized")
	}
	return session.(*model.User), nil
}
