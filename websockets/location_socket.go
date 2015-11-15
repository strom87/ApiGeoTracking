package websockets

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/strom87/ApiGeoTracking/core/logger"
	"github.com/strom87/ApiGeoTracking/models"
)

// LocationSocket struct
type LocationSocket struct {
	logger      *logger.Logger
	connections map[*websocket.Conn]string
}

// NewLocationSocket pointer of LocationSocket
func NewLocationSocket() *LocationSocket {
	location := &LocationSocket{logger: logger.NewLogger(), connections: map[*websocket.Conn]string{}}
	location.startPing()
	return location
}

// ServeHTTP handles web requests
func (s LocationSocket) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Log("websocket connection open error:", err)
	}
	defer conn.Close()

	id := mux.Vars(r)["id"]
	s.addConnection(id, conn)

	for {
		model := models.GeoLocationModel{}
		if err := websocket.ReadJSON(conn, &model); err != nil {
			s.logger.Log("Conn read message error:", err)
			s.disconnect(conn)
			return
		}

		s.sendLocationMessage(model, s.connections[conn])
	}
}

func (s LocationSocket) addConnection(id string, conn *websocket.Conn) {
	if _, ok := s.connections[conn]; !ok {
		s.connections[conn] = id
	}
}

func (s LocationSocket) sendLocationMessage(model models.GeoLocationModel, id string) {
	model.ID = id
	for conn := range s.connections {
		if err := conn.WriteJSON(model); err != nil {
			s.disconnect(conn)
		}
	}
}

func (s LocationSocket) disconnect(conn *websocket.Conn) {
	model := models.GeoLocationModel{ID: s.connections[conn], Disconnected: true}
	delete(s.connections, conn)

	for conn := range s.connections {
		if err := conn.WriteJSON(model); err != nil {
			s.disconnect(conn)
		}
	}
}

func (s LocationSocket) startPing() {
	go func() {
		ticker := time.NewTicker(pingInterval)

		for {
			select {
			case <-ticker.C:
				for conn := range s.connections {
					if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
						s.disconnect(conn)
					}
				}
			}
		}
	}()
}
