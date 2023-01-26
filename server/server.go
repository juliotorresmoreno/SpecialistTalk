package server

import (
	"os"
	"strings"
	"time"

	"github.com/juliotorresmoreno/freelive/configs"
	"github.com/juliotorresmoreno/freelive/handler"
	middleware_app "github.com/juliotorresmoreno/freelive/middleware"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerHTTP struct {
	*echo.Echo
}

func (s *ServerHTTP) Listen() error {
	conf := configs.GetConfig()
	host := conf.Host + ":" + conf.Port
	return s.Start(host)
}

func NewServer() *ServerHTTP {
	e := echo.New()
	svr := &ServerHTTP{e}

	conf := configs.GetConfig()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path != "ws"
		},
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "metrics")
		},
	}))
	e.Use(middleware.CORS())
	// e.Use(middleware.CSRF())

	handler.AttachWS(e.Group("/ws", middleware_app.Session))

	handler.AttachSwaggerApi(e.Group("/docs"))

	api := e.Group("/api/v1", middleware_app.Session)
	handler.AttachAuth(api.Group("/auth"))
	handler.AttachUsers(api.Group("/users"))

	static := e.Group("/*")
	if conf.Env == "development" {
		static.Use(middleware_app.NoCache)
	}
	static.GET("", func(c echo.Context) error {
		_, err := os.Lstat("website/dist/" + c.Request().URL.Path)
		if err != nil {
			return c.File("website/dist/index.html")
		}
		return c.File("website/dist/" + c.Request().URL.Path)
	})

	return svr
}
