package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/juliotorresmoreno/freelive/model"
	"github.com/labstack/echo/v4"
)

type client struct {
	conn     *websocket.Conn
	username string
}

type Message struct {
	Username     string
	Notification model.Notification
}

var register = make(chan *client)
var remove = make(chan *client)
var SendToClient = make(chan *Message)

var clients = map[string]map[*client]bool{}

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
			Type:    "error",
			Message: "Unauthorized"},
		)
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*model.User)

	cli := &client{conn: ws, username: session.Username}
	register <- cli

	SendToClient <- &Message{
		session.Username,
		model.Notification{
			Type:    string(model.NotificationEvent),
			Message: "Connected",
		},
	}

	readWS(cli)
	return nil
}

func (u *HandlerWS) Register() {
	for {
		select {
		case c := <-register:
			username := c.username

			if _, ok := clients[username]; !ok {
				clients[username] = make(map[*client]bool)
			}
			clients[username][c] = true
		case c := <-remove:
			username := c.username

			if slot, ok := clients[username]; ok {
				delete(slot, c)
			}
		case c := <-SendToClient:
			if slot, ok := clients[c.Username]; ok {
				for client := range slot {
					_ = client.conn.WriteJSON(c.Notification)
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
			remove <- cli
			_ = cli.conn.Close()
		}
	}
}
