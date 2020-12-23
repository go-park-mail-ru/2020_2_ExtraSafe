package websocket

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type Hub interface {
	CheckUser(sessionID string) bool
	Broadcast(message interface{})
	Send(message models.NotificationMessage)
	Register(client *Client)
	Unregister(client *Client)
}
