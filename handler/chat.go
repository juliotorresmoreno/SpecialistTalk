package handler

import (
	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/labstack/echo/v4"
)

type ChatsHandler struct{}

// AttachChats s
func AttachChats(g *echo.Group) {
	u := &ChatsHandler{}
	g.GET("", u.findUsers)
	g.GET("/:user_id", u.findUser)
	g.PUT("", u.addUser)
	g.PATCH("/:user_id", u.updateUser)
}

func (that *ChatsHandler) findUser(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}

	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return err
	}

	u := &model.User{}
	_, err = conn.SessionWithACL().Where("id = ?", c.Param("user_id")).Get(u)
	if err != nil {
		return echo.NewHTTPError(501, helper.ParseError(err).Error())
	}
	return c.JSON(200, u)
}

func (that *ChatsHandler) findUsers(c echo.Context) error {
	return nil
}

func (that *ChatsHandler) addUser(c echo.Context) error {
	return nil
}

func (that *ChatsHandler) updateUser(c echo.Context) error {
	return nil
}
