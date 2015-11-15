package websockets

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pingInterval = time.Second * 60
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
