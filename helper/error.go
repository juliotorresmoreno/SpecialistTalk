package helper

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// ParseError s
func ParseError(err error) error {
	v := err.Error()
	if len(v) >= 17 && v[:17] == "pq: duplicate key" {
		s := strings.Split(v, "\"")
		s = strings.Split(s[1], "_")
		f := strings.Join(s[2:], "_")
		return fmt.Errorf("%v: %v already exists", f, f)
	}
	if len(v) >= 8 && v[:8] == "dial tcp" {
		return errors.New("database is not running")
	}
	errString := strings.ToLower(err.Error())
	return errors.New(errString)
}

func MakeHTTPError(status int, err error) *echo.HTTPError {
	return &echo.HTTPError{
		Code:    status,
		Message: err.Error(),
	}
}

var HTTPStatusNotContent = echo.NewHTTPError(http.StatusOK, "no content")

var HTTPErrorNotImplementedError = echo.NewHTTPError(http.StatusNotImplemented, "not implemented")
var HTTPErrorInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
var HTTPErrorUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
var HTTPErrorNotFound = echo.NewHTTPError(http.StatusNotFound, "not found")
var HTTPErrorBadRequest = echo.NewHTTPError(http.StatusNotFound, "bad request")
