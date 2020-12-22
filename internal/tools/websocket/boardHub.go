package websocket

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"sync"
)

type BoardHub struct {
	boardID int64
	users map[string]*Client
	broadcast chan interface{}
	register chan *Client
	unregister chan *Client
	stop chan bool
	hubs *sync.Map
}

func NewHub(boardID int64, hubs *sync.Map) *BoardHub {
	return &BoardHub{
		boardID: boardID,
		broadcast:  make(chan interface{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		users: make(map[string]*Client),
		stop: make(chan bool),
		hubs: hubs,
	}
}

func (h *BoardHub) Run() {
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
				h.hubs.Delete(h.boardID)
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

func (h *BoardHub) StopHub() {
	h.stop <- true
}

func (h *BoardHub) CheckUser(sessionID string) bool {
	if _, ok := h.users[sessionID]; ok {
		return true
	}
	return false
}

func (h *BoardHub) Broadcast(message interface{}) {
	h.broadcast <- message
}

func (h * BoardHub) Register(client *Client) {
	h.register <- client
}

func (h * BoardHub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *BoardHub) Send(message models.NotificationMessage) {
	h.broadcast <- message
}

func (h *BoardHub) GetClients() {
	fmt.Println(h.users)
}