package websocket

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Hub struct {
	boardID int64
	users map[string]*Client
	broadcast chan interface{}
	register chan *Client
	unregister chan *Client
	stop chan bool
	hubs *map[int64]*Hub
}

func NewHub(boardID int64, hubs *map[int64]*Hub) *Hub {
	return &Hub{
		boardID: boardID,
		broadcast:  make(chan interface{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		users: make(map[string]*Client),
		stop: make(chan bool),
		hubs: hubs,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.users[client.sessionID] = client

		case client := <-h.unregister:
			if _, ok := h.users[client.sessionID]; ok {
				delete(h.users, client.sessionID)
				close(client.send)
			}
			if len(h.users) == 0 {
				delete(*h.hubs, h.boardID)
				return
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

func (h *Hub) StopHub() {
	h.stop <- true
}

func (h * Hub) CheckUser(sessionID string) bool {
	if _, ok := h.users[sessionID]; ok {
		return true
	}
	return false
}

func (h * Hub) Broadcast(message interface{}) {
	h.broadcast <- message
}

func (h * Hub) GetClients() {
	fmt.Println(h.users)
}