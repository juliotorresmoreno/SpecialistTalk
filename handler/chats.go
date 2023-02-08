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

type findUserResponse struct {
	Name string      `json:"name"`
	Code string      `json:"code"`
	User *model.User `json:"user"`
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
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}
	if !ok {
		token := helper.GenerateToken()
		chat.Code = token
		chat.Name = u.Name + " " + u.LastName
		chat.ACL = &model.ACL{Owner: session.Username}
		chat.Status = model.ChatStatusActive
		if err = chat.Check(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
		}
		_, _ = conn.InsertOne(chat)

		chat2 := &model.Chat{UserID: session.ID}
		chat2.Code = token
		chat2.Name = session.Name + " " + session.LastName
		chat2.ACL = &model.ACL{Owner: u.Username}
		chat.Status = model.ChatStatusCreated
		if err = chat.Check(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
		}
		_, _ = conn.InsertOne(chat)
	}

	return c.JSON(200, &findUserResponse{
		Name: chat.Name,
		Code: chat.Code,
		User: u,
	})
}

func (that *ChatsHandler) findUsers(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}

	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}

	chats := make([]model.Chat, 0)
	conn.ShowSQL(true)
	err = conn.Find(&chats)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}

	return c.JSON(200, chats)
}

func (that *ChatsHandler) addUser(c echo.Context) error {
	return nil
}

func (that *ChatsHandler) updateUser(c echo.Context) error {
	return nil
}
