package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 开启压缩
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(ctx *gin.Context) {
	c, e := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if e != nil {
		ctx.JSON(500, "server error")
		return
	}

	client := &Client{
		hub:  hub,
		conn: c,
		send: make(chan []byte, 256),
	}

	client.hub.register <- client

	go client.writePump()
}
