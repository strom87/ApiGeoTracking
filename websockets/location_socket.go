package websockets

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/strom87/ApiGeoTracking/core/factory"
	"github.com/strom87/ApiGeoTracking/core/logger"
	"github.com/strom87/ApiGeoTracking/models"
)

// LocationInfo struct
type LocationInfo struct {
	ID   string
	Name string
}

// LocationSocket struct
type LocationSocket struct {
	logger      *logger.Logger
	connections map[*websocket.Conn]LocationInfo
}

// NewLocationSocket pointer of LocationSocket
func NewLocationSocket() *LocationSocket {
	location := &LocationSocket{logger: logger.NewLogger(), connections: map[*websocket.Conn]LocationInfo{}}
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

	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	id := mux.Vars(r)["id"]
	if id == "" || !bson.IsObjectIdHex(id) {
		conn.Close()
		s.logger.Log("LocationSocket ServeHTTP id is not a bson id.", err)
		json.NewEncoder(w).Encode(models.Response{ErrorCode: 5})
		return
	}

	dbConn := factory.NewDatabase().Connection()
	if err := dbConn.Open(); err != nil {
		s.logger.Log("LocationSocket ServeHTTP could not open db connection.", err)
		json.NewEncoder(w).Encode(models.Response{ErrorCode: 2})
		return
	}
	defer dbConn.Close()

	user, err := factory.NewDatabase().User(dbConn).FindByID(bson.ObjectIdHex(id))
	if err != nil || user == nil {
		s.logger.Log("LocationSocket ServeHTTP find by id.", err)
		json.NewEncoder(w).Encode(models.Response{ErrorCode: 3})
		return
	}

	s.addConnection(LocationInfo{ID: id, Name: user.Name}, conn)

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

func (s LocationSocket) addConnection(loc LocationInfo, conn *websocket.Conn) {
	if _, ok := s.connections[conn]; !ok {
		s.connections[conn] = loc
	}
}

func (s LocationSocket) sendLocationMessage(model models.GeoLocationModel, loc LocationInfo) {
	model.ID = loc.ID
	model.Name = loc.Name
	for conn := range s.connections {
		if err := conn.WriteJSON(model); err != nil {
			s.disconnect(conn)
		}
	}
}

func (s LocationSocket) disconnect(conn *websocket.Conn) {
	model := models.GeoLocationModel{ID: s.connections[conn].ID, Disconnected: true}
	delete(s.connections, conn)

	for conn := range s.connections {
		if err := conn.WriteJSON(model); err != nil {
			s.disconnect(conn)
		}
	}
}

func (s LocationSocket) startPing() {
	go func() {
		ticker := time.NewTicker(pingPeriod)

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
