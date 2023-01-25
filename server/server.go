package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/juliotorresmoreno/freelive/configs"
	"github.com/juliotorresmoreno/freelive/handler"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerHTTP struct {
	*echo.Echo
}

// Listen s
func (s *ServerHTTP) Listen() error {
	conf := configs.GetConfig()
	host := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	return s.Start(host)
}

func NewServer() *ServerHTTP {
	e := echo.New()
	svr := &ServerHTTP{e}

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "metrics")
		},
	}))
	e.Use(middleware.CORS())
	// e.Use(middleware.CSRF())

	api := e.Group("/api/v1")
	handler.AttachSwaggerApi(e.Group("/docs"))
	handler.AttachAuth(api.Group("/auth"))

	handler.AttachUsers(api.Group("/users"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return svr
}
