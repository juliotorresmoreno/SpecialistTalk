package middleware

import (
	"strconv"
	"time"

	"github.com/juliotorresmoreno/freelive/db"
	"github.com/juliotorresmoreno/freelive/model"
	"github.com/juliotorresmoreno/freelive/services"
	"github.com/labstack/echo/v4"
)

func Session(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := ""
		cookie, _ := c.Cookie("token")
		if cookie != nil {
			token = cookie.Value
		} else {
			token = c.QueryParams().Get("token")
		}
		if token == "" {
			return handler(c)
		}
		redisCli := services.NewRedis()
		go redisCli.Close()
		r := redisCli.Get(token).Val()
		if r != "" {
			conn, err := db.GetConnectionPool()
			if err == nil {
				id, _ := strconv.Atoi(r)
				u := &model.User{ID: id}
				ok, _ := conn.Select("id, username, name, lastname, acl").Get(u)
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