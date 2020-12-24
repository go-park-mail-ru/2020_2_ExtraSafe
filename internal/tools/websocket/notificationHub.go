package websocket

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type NotificationHub struct {
	users map[string]*Client
	broadcast chan interface{}
	notification chan models.NotificationMessage
	register chan *Client
	unregister chan *Client
	stop chan bool
	hubs *map[int64]*NotificationHub
}

func NewNotificationHub() *NotificationHub {
	return &NotificationHub{
		broadcast:  make(chan interface{}),
		notification: make(chan models.NotificationMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		users: make(map[string]*Client),
		stop: make(chan bool),
	}
}

func (h *NotificationHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.users[client.sessionID] = client

		case client := <-h.unregister:
			if _, ok := h.users[client.sessionID]; ok {
				delete(h.users, client.sessionID)
				close(client.send)
			}

		case message := <- h.notification:
			for id, user := range h.users {
				if user.userID == message.UserID {
					select {
					case user.send <- message:
					default:
						close(user.send)
						delete(h.users, id)
					}
				}
			}

		case message := <-h.broadcast:
			for id, user := range h.users {
				if id == message.(models.WS).SessionID {
					continue
				}
				select {
				case user.send <- message:
				default:
					close(user.send)
					delete(h.users, id)
				}
			}

		case status := <- h.stop:
			if status == true {
				return
			}
		}
	}
}

func (h *NotificationHub) StopHub() {
	h.stop <- true
}

func (h * NotificationHub) CheckUser(sessionID string) bool {
	if _, ok := h.users[sessionID]; ok {
		return true
	}
	return false
}

func (h * NotificationHub) Broadcast(message interface{}) {
	h.broadcast <- message
}

func (h * NotificationHub) Register(client *Client) {
	h.register <- client
}

func (h * NotificationHub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *NotificationHub) Send(message models.NotificationMessage) {
	h.notification <- message
}

func (h * NotificationHub) GetClients() {
	fmt.Println(h.users)
}

