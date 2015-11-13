package factory

import "github.com/strom87/ApiGeoTracking/db"

const (
	dbName       = "geo-tracking"
	dbConnection = "mongodb://localhost:27017"
)

// Database struct
type Database struct{}

// NewDatabase pointer to Database
func NewDatabase() *Database {
	return &Database{}
}

// Connection generates a connection
func (d Database) Connection() *db.Connection {
	return db.NewConnection(dbName, dbConnection)
}

// User returns pointer of UserProvider
func (d Database) User(connection *db.Connection) *db.UserProvider {
	return db.NewUserProvider(connection)
}
