package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

// ParseError s
func ParseError(err error) error {
	v := err.Error()
	if v[:17] == "pq: duplicate key" {
		s := strings.Split(v, "\"")
		s = strings.Split(s[1], "_")
		f := strings.Join(s[2:], "_")
		return fmt.Errorf("%v: %v already exists", f, f)
	}
	if v[:8] == "dial tcp" {
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
