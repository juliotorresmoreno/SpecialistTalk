package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/juliotorresmoreno/SpecialistTalk/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessagesHandler struct{}

// AttachMessages s
func AttachMessages(g *echo.Group) {
	u := &MessagesHandler{}
	g.GET("", u.find)
	g.GET("/:code", u.get)
	g.POST("", u.add)
	g.PATCH("/:user_id", u.update)
}

func (that *MessagesHandler) get(c echo.Context) error {
	_, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}

	return helper.HTTPErrorNotImplementedError
}

func (that *MessagesHandler) find(c echo.Context) error {
	_, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}

	return helper.HTTPErrorNotImplementedError
}

type Attachment struct {
	Name string `json:"name" valid:"required"`
	Body string `json:"body" valid:"required"`
}

func (that *MessagesHandler) updateNotificationChat(code string, username string) error {
	conn, err := db.GetConnectionPool()
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	chats := make([]model.Chat, 0)
	err = conn.Where("code = ? and owner <> ?", code, username).Find(&chats)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	for _, chat := range chats {
		chat.Notifications += 1
		_, _ = conn.Cols("notifications").
			Where("id = ?", chat.ID).
			Update(chat)
		dispatchContactsUpdate(chat.Owner)
	}

	return nil
}

type POSTMessagesAddPayload struct {
	Code        string       `json:"code"        valid:"required"`
	Message     string       `json:"message"     valid:"required"`
	Attachments []Attachment `json:"attachments"`
}

func (that *MessagesHandler) add(c echo.Context) error {
	payload := &POSTMessagesAddPayload{}
	session, err := helper.GetPayload(c, payload)
	if err != nil {
		return err
	}

	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	chat := &model.Chat{Code: payload.Code}
	ok, err := conn.Get(chat)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	} else if !ok {
		return helper.HTTPErrorNotFound
	}
	if err = that.updateNotificationChat(chat.Code, session.Username); err != nil {
		return err
	}

	mongoCli, err := services.GetPoolMongo()
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	collectionName := "conversation_" + payload.Code
	collection := mongoCli.Database(conf.Mongo.ConversationDB).Collection(collectionName)

	id := primitive.NewObjectID()
	now := time.Now()
	data := &model.Message{
		ID:        &id,
		Message:   payload.Message,
		From:      session.ID,
		FromName:  session.Name + " " + session.LastName,
		CreatedAt: &now,
	}
	_, err = collection.InsertOne(context.Background(), data)
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	dispatchMessageToGroup(payload.Code, data)

	return helper.HTTPStatusNotContent
}

func (that *MessagesHandler) update(c echo.Context) error {
	_, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}

	return helper.HTTPErrorNotImplementedError
}
