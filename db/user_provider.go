package db

import "gopkg.in/mgo.v2/bson"

// UserProvider connection to user document
type UserProvider struct {
	*Connection
	Collection string
}

// NewUserProvider pointer to user provider
func NewUserProvider(connection *Connection) *UserProvider {
	return &UserProvider{Connection: connection, Collection: "users"}
}

// FindByID get one user by id
func (p UserProvider) FindByID(id bson.ObjectId) (*User, error) {
	db := p.Connection.GetCollection(p.Collection)

	user := &User{}
	if err := db.Find(bson.M{"_id": id}).One(user); err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail get one user by email
func (p UserProvider) FindByEmail(email string) (*User, error) {
	db := p.Connection.GetCollection(p.Collection)

	user := &User{}
	if err := db.Find(bson.M{"email": email}).One(user); err != nil {
		return nil, err
	}

	return user, nil
}
