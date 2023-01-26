package model

type NotificationType string

var NotificationError NotificationType = "error"
var NotificationMessage NotificationType = "message"
var NotificationEvent NotificationType = "event"

type Notification struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}
