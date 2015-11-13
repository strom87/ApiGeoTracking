package db

import "gopkg.in/mgo.v2"

// Connection db info
type Connection struct {
	DbName           string
	ConnectionString string
	Session          *mgo.Session
}

// NewConnection returns pointer to Connection struct
func NewConnection(dbName, connectionString string) *Connection {
	return &Connection{DbName: dbName, ConnectionString: connectionString}
}

// Open open up a active session to mongodb
func (c *Connection) Open() error {
	session, err := mgo.Dial(c.ConnectionString)
	if err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	c.Session = session

	return nil
}

// Close close the active session
func (c *Connection) Close() {
	c.Session.Close()
}

// GetCollection get one collection from the current session
func (c Connection) GetCollection(collection string) *mgo.Collection {
	return c.Session.DB(c.DbName).C(collection)
}
