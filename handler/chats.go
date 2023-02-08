package handler

import (
	"net/http"
	"strconv"

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

	userID, _ := strconv.Atoi(c.Param("user_id"))
	if session.ID == userID {
		return helper.HTTPErrorUnauthorized
	}

	conn, err := db.GetConnectionPool()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}

	u := &model.User{}
	ok, err := conn.Where("id = ?", c.Param("user_id")).Get(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}
	if !ok {
		return helper.HTTPErrorNotFound
	}

	conf := configs.GetConfig()
	conn, err = db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}

	chat := &model.Chat{UserID: u.ID}
	ok, err = conn.Get(chat)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	if !ok {
		chat.Code = "any"
		chat.Name = "any"
		chat.ACL = &model.ACL{Owner: session.Username}
		chat.UserID = userID
		_, _ = conn.InsertOne(chat)
	}

	return c.JSON(200, chat)
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
