package helper

import (
	"strconv"

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
