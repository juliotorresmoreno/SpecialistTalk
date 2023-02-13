package helper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// ParseError s
func ParseError(err interface{}) string {
	v := ""
	switch result := err.(type) {
	case []byte:
		v = string(result)
	case string:
		v = result
	case error:
		v = result.Error()
	default:
		v = fmt.Sprintf("%v", err)
	}

	if len(v) >= 17 && v[:17] == "pq: duplicate key" {
		s := strings.Split(v, "\"")
		s = strings.Split(s[1], "_")
		f := strings.Join(s[2:], "_")
		return fmt.Sprintf("%v: %v already exists", f, f)
	}
	if len(v) >= 8 && v[:8] == "dial tcp" {
		return "database is not running"
	}

	return strings.ToLower(v)
}

func MakeHTTPError(status int, err interface{}) *echo.HTTPError {
	return &echo.HTTPError{
		Code:    status,
		Message: ParseError(err),
	}
}

var HTTPStatusNotContent = echo.NewHTTPError(http.StatusOK, "no content")

var HTTPErrorNotImplementedError = echo.NewHTTPError(http.StatusNotImplemented, "not implemented")
var HTTPErrorInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
var HTTPErrorUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
var HTTPErrorNotFound = echo.NewHTTPError(http.StatusNotFound, "not found")
var HTTPErrorBadRequest = echo.NewHTTPError(http.StatusNotFound, "bad request")
