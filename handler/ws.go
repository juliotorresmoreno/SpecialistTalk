package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/labstack/echo/v4"
)

type client struct {
	conn     *websocket.Conn
	username string
}

type MessageToClient struct {
	Username     string
	Notification *model.Notification
}
type MessageToGroup struct {
	Code         string
	Notification *model.Notification
}

var register = make(chan *client)
var removeClient = make(chan *client)
var removeUser = make(chan string)
var SendToClient = make(chan *MessageToClient)
var SendToGroup = make(chan *MessageToGroup)

var clients = map[string]map[*client]bool{}

func dispatchContactsUpdate(username string) {
	conn, err := db.GetConnectionPool()
	if err != nil {
		return
	}

	chats := make([]model.Chat, 0)
	err = conn.Where("owner = ?", username).Find(&chats)
	if err != nil {
		return
	}

	SendToClient <- &MessageToClient{
		Username: username,
		Notification: &model.Notification{
			Type:    "contacts_update",
			Payload: chats,
		},
	}
}

func dispatchMessageToGroup(code string, data interface{}) {
	SendToGroup <- &MessageToGroup{
		Code: code,
		Notification: &model.Notification{
			Type: "message",
			Payload: map[string]interface{}{
				"code": code,
				"data": data,
			},
		},
	}
}

func dispatchDisconnect(username string) {
	removeUser <- username
}

type HandlerWS struct {
	Upgrader *websocket.Upgrader
}

func AttachWS(g *echo.Group) {
	u := &HandlerWS{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	go u.Register()

	g.GET("", u.GET)
}

func (u *HandlerWS) GET(c echo.Context) error {
	ws, err := u.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	_session := c.Get("session")
	if _session == nil {
		_ = ws.WriteJSON(model.Notification{
			Type:    string(model.NotificationError),
			Payload: "Unauthorized",
		})
		return helper.HTTPErrorUnauthorized
	}
	session := _session.(*model.User)

	cli := &client{conn: ws, username: session.Username}
	register <- cli

	SendToClient <- &MessageToClient{
		session.Username,
		&model.Notification{
			Type:    string(model.NotificationEvent),
			Payload: "Connected",
		},
	}

	readWS(cli)
	return nil
}

func (u *HandlerWS) Register() {
	conn, err := db.GetConnectionPool()
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case c := <-register:
			username := c.username

			if _, ok := clients[username]; !ok {
				clients[username] = make(map[*client]bool)
			}
			clients[username][c] = true
		case c := <-removeUser:
			username := c

			if slot, ok := clients[username]; ok {
				for client := range slot {
					_ = client.conn.Close()
				}
				delete(clients, username)
			}
		case c := <-removeClient:
			username := c.username

			if slot, ok := clients[username]; ok {
				_ = c.conn.Close()
				delete(slot, c)
			}
		case c := <-SendToClient:
			username := c.Username
			if slot, ok := clients[username]; ok {
				for client := range slot {
					_ = client.conn.WriteJSON(c.Notification)
				}
			}
		case c := <-SendToGroup:
			chats := make([]model.Chat, 0)
			err = conn.Where("code = ?", c.Code).Find(&chats)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, chat := range chats {
				username := chat.Owner
				if slot, ok := clients[username]; ok {
					for client := range slot {
						_ = client.conn.WriteJSON(c.Notification)
					}
				}
			}
		}
	}
}

// El socket suele cambiar su estado a lo largo del tiempo por ese motivo no suelo
// leer los mensajes por este canal ya que me tocaria gestionar correctamente el estado
// del socket a la hora de enviar notificaciones. En cambio prefiero usar endpoints normales
// y usar este medio en una sola direccion, es decir hacia el cliente.
// Esta funcion solo existe para que el socket no se cierre.
func readWS(cli *client) {
	for {
		_, _, err := cli.conn.ReadMessage()
		if err != nil {
			removeClient <- cli
			_ = cli.conn.Close()
			break
		}
	}
}
