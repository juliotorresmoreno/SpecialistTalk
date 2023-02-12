package handler

import (
	"context"
	"net/http"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/juliotorresmoreno/SpecialistTalk/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatsHandler struct{}

// AttachChats s
func AttachChats(g *echo.Group) {
	u := &ChatsHandler{}
	g.GET("", u.find)
	g.GET("/:code", u.get)
	g.POST("", u.add)
	g.PATCH("/:user_id", u.update)
}

func (that *ChatsHandler) get(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}
	code := c.Param("code")

	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}

	chat := &model.Chat{Code: code}
	if ok, err := conn.Get(chat); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	} else if !ok {
		return helper.HTTPErrorUnauthorized
	}

	mongoCli, err := services.GetPoolMongo()
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	collectionName := "conversation_" + code
	collection := mongoCli.Database(conf.Mongo.Database).Collection(collectionName)
	limit := int64(10)
	curr, _ := collection.Find(context.Background(), map[string]interface{}{}, &options.FindOptions{
		Limit: &limit,
		Sort:  bson.D{{Key: "created_at", Value: -1}},
	})

	messages := make([]model.Message, 0)
	_ = curr.All(context.Background(), &messages)
	helper.Reverse(messages)

	return c.JSON(200, map[string]interface{}{
		"id":       chat.ID,
		"user_id":  chat.UserID,
		"name":     chat.Name,
		"code":     chat.Code,
		"status":   chat.Status,
		"messages": messages,
	})
}

func (that *ChatsHandler) find(c echo.Context) error {
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
	err = conn.Find(&chats)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}

	return c.JSON(200, chats)
}

type POSTChatsAddPayload struct {
	UserID int `json:"user_id" valid:"required"`
}

func (that *ChatsHandler) add(c echo.Context) error {
	payload := &POSTChatsAddPayload{}
	session, err := helper.GetPayload(c, payload)
	if err != nil {
		return err
	}

	if session.ID == payload.UserID {
		return helper.HTTPErrorUnauthorized
	}

	conn, err := db.GetConnectionPool()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}

	u := &model.User{}
	ok, err := conn.Where("id = ?", payload.UserID).Get(u)
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
		chat2.Status = model.ChatStatusCreated
		if err = chat.Check(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
		}
		_, _ = conn.InsertOne(chat2)
	}

	SendToClient <- &MessageToClient{
		Username: session.Username,
		Notification: &model.Notification{
			Type: "contacts_update",
		},
	}

	return c.JSON(200, chat)
}

func (that *ChatsHandler) update(c echo.Context) error {
	return nil
}
