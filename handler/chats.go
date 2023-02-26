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
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	chat := &model.Chat{Code: code}
	if ok, err := conn.Get(chat); err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	} else if !ok {
		return helper.HTTPErrorUnauthorized
	}

	chat.Notifications = 0
	_, _ = conn.Cols("notifications").
		Where("id = ?", chat.ID).
		Update(chat)
	dispatchContactsUpdate(chat.Owner)

	mongoCli, err := services.GetPoolMongo()
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	collectionName := "conversation_" + code
	collection := mongoCli.Database(conf.Mongo.ConversationDB).Collection(collectionName)
	limit := int64(10)
	curr, _ := collection.Find(context.Background(), map[string]interface{}{}, &options.FindOptions{
		Limit: &limit,
		Sort:  bson.D{{Key: "created_at", Value: -1}},
	})

	messages := make([]model.Message, 0)
	_ = curr.All(context.Background(), &messages)
	messages = helper.Reverse(messages)

	user := &model.User{Username: chat.ToUserName}
	if _, err = conn.NewSessionFree().Get(user); err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(200, map[string]interface{}{
		"id":       chat.ID,
		"user":     user,
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
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	chats := make([]model.Chat, 0)
	err = conn.Find(&chats)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
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

	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	u := &model.User{ID: payload.UserID}
	ok, err := conn.NewSessionFree().Get(u)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	if !ok {
		return helper.HTTPErrorNotFound
	}

	chat := &model.Chat{ToUserName: u.Username}
	ok, err = conn.Get(chat)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	if !ok {
		token := helper.GenerateToken()
		chat.Code = token
		chat.Name = u.Name + " " + u.LastName
		chat.Owner = session.Username
		chat.ToUserName = u.Username
		chat.Status = model.ChatStatusActive
		if err = chat.Check(); err != nil {
			return helper.MakeHTTPError(http.StatusInternalServerError, err)
		}
		_, _ = conn.InsertOne(chat)

		other_chat := &model.Chat{}
		other_chat.Code = token
		other_chat.Name = session.Name + " " + session.LastName
		other_chat.Owner = u.Username
		other_chat.ToUserName = session.Username
		other_chat.Status = model.ChatStatusCreated
		if err = chat.Check(); err != nil {
			return helper.MakeHTTPError(http.StatusInternalServerError, err)
		}
		_, _ = conn.NewSessionFree().InsertOne(other_chat)
	} else if chat.Status == model.ChatStatusCreated {
		chat.Status = model.ChatStatusActive
		if _, err = conn.Update(chat); err != nil {
			return helper.MakeHTTPError(http.StatusInternalServerError, err)
		}
	}

	dispatchContactsUpdate(session.Username)

	return c.JSON(200, chat)
}

func (that *ChatsHandler) update(c echo.Context) error {
	return helper.HTTPErrorNotImplementedError
}
