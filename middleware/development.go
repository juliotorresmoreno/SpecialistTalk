package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

type TimeSleepConfig struct {
	Duration time.Duration
}

func TimeSleep(config *TimeSleepConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			time.Sleep(config.Duration)
			return next(c)
		}
	}
}
