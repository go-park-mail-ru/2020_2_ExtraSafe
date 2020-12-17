package websocket

type Hub struct {
	boardID int64
	users map[int64]*Client
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
		users: make(map[int64]*Client),
		stop: make(chan bool),
		hubs: hubs,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.users[client.ID] = client

		case client := <-h.unregister:
			if _, ok := h.users[client.ID]; ok {
				delete(h.users, client.ID)
				close(client.send)
			}
			if len(h.users) == 0 {
				delete(*h.hubs, h.boardID)
				return
			}

		case message := <-h.broadcast:
			for id, user := range h.users {
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

func (h * Hub) CheckUser(userID int64) bool {
	if _, ok := h.users[userID]; ok {
		return true
	}
	return false
}

func (h * Hub) Broadcast(message interface{}) {
	h.broadcast <- message
}