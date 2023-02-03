package handler

import (
	"net/url"
	"os"
	"os/exec"
	"path"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	middleware_app "github.com/juliotorresmoreno/SpecialistTalk/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type StaticHandler struct{}

// AuthHandler s
func AttachStatic(g *echo.Group) {
	var cmd *exec.Cmd
	conf := configs.GetConfig()

	if conf.Env != "production" {
		cmd = exec.Command("npm", "run", "dev")
	} else {
		cmd = exec.Command("npm", "run", "build")

	}
	cmd.Dir = path.Join(cmd.Dir, "website")
	cmd.Stderr = os.Stderr
	go func() {
		_ = cmd.Run()
	}()

	if conf.Env == "production" {
		g.GET("", func(c echo.Context) error {
			_, err := os.Lstat("website/dist/" + c.Request().URL.Path)
			if err != nil {
				return c.File("website/dist/index.html")
			}
			return c.File("website/dist/" + c.Request().URL.Path)
		})
	} else {
		g.Use(middleware_app.NoCache)
		g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: &url.URL{
					Host:   "localhost:19000",
					Path:   "/",
					Scheme: "http",
				},
			},
		})))
	}
}
