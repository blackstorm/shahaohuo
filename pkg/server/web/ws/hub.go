package ws

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/task"
)

type clientHub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

var hub *clientHub

func (h *clientHub) run() {
	task.AddListener(h)
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *clientHub) OnChange(hs []orm.BusinessHaohuo) {
	if str, e := json.Marshal(&hs); e == nil {
		h.broadcast <- str
	} else {
		logrus.Error(e)
	}
}

func InitHub() {
	hub = &clientHub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}

	go hub.run()
}
