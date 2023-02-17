package middleware

import (
	"strconv"
	"time"

	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/juliotorresmoreno/SpecialistTalk/services"
	"github.com/labstack/echo/v4"
)

func Session(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("X-API-Key")
		if token == "" {
			token = c.QueryParams().Get("token")
		}
		if token == "" {
			return handler(c)
		}
		redisCli := services.GetPoolRedis()
		r := redisCli.Get(token).Val()
		if r != "" {
			conn, err := db.GetConnectionPool()
			if err == nil {
				id, _ := strconv.Atoi(r)
				u := &model.User{ID: id}
				ok, _ := conn.Select("id, username, name, lastname, owner").Get(u)
				if ok {
					u.Password = ""
					u.ValidPassword = ""
					c.Set("session", u)
					redisCli.Set(token, r, 24*time.Hour)
				}
			}
		}
		return handler(c)
	}
}
