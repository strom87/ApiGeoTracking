package websockets

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/strom87/ApiGeoTracking/core/logger"
	"github.com/strom87/ApiGeoTracking/models"
)

// LocationSocket struct
type LocationSocket struct {
	logger      *logger.Logger
	connections map[string]*websocket.Conn
}

// NewLocationSocket pointer of LocationSocket
func NewLocationSocket() *LocationSocket {
	return &LocationSocket{logger: logger.NewLogger(), connections: map[string]*websocket.Conn{}}
}

// ServeHTTP handles web requests
func (s LocationSocket) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Log("websocket connection open error:", err)
	}
	defer conn.Close()

	for {
		model := models.GeoLocationModel{}
		if err := websocket.ReadJSON(conn, &model); err != nil {
			s.logger.Log("Conn read message error:", err)
			return
		}

		s.addConnection(model.ID, conn)

		if model.Long == 0.0 && model.Lat == 0.0 {
			continue
		}

		s.sendLocationMessage(model)
	}
}

func (s LocationSocket) addConnection(id string, conn *websocket.Conn) {
	if _, ok := s.connections[id]; !ok {
		s.connections[id] = conn
	}
}

func (s LocationSocket) sendLocationMessage(v interface{}) {
	for key, conn := range s.connections {
		if err := conn.WriteJSON(v); err != nil {
			delete(s.connections, key)
			s.sendDisconnectedMessage(key)
		}
	}
}

func (s LocationSocket) sendDisconnectedMessage(id string) {
	model := models.GeoLocationModel{ID: id, Disconnected: true}

	for key, conn := range s.connections {
		if err := conn.WriteJSON(model); err != nil {
			delete(s.connections, key)
			s.sendDisconnectedMessage(key)
		}
	}
}
