package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationType string

var NotificationError NotificationType = "error"
var NotificationMessage NotificationType = "message"
var NotificationEvent NotificationType = "event"

type Notification struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
type Message struct {
	ID        *primitive.ObjectID `json:"id"         bson:"_id"`
	Message   string              `json:"message"    bson:"message"`
	From      int                 `json:"from"       bson:"from"`
	FromName  string              `json:"from_name"  bson:"from_name"`
	CreatedAt *time.Time          `json:"created_at" bson:"created_at"`
}
