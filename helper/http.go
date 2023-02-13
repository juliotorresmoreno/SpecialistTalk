package helper

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/asaskevich/govalidator"
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
		return nil, HTTPErrorUnauthorized
	}
	return session.(*model.User), nil
}

func GetPayload(c echo.Context, payload interface{}) (*model.User, error) {
	session, err := ValidateSession(c)
	if err != nil {
		return session, err
	}

	kindOfJ := reflect.ValueOf(payload).Kind()
	if kindOfJ != reflect.Ptr {
		return session, echo.NewHTTPError(http.StatusInternalServerError, "payload must a pointer")
	}

	err = c.Bind(payload)
	if err != nil {
		return session, HTTPErrorBadRequest
	}

	_, err = govalidator.ValidateStruct(payload)
	if err != nil {
		return session, MakeHTTPError(http.StatusBadRequest, err)
	}

	return session, nil
}
