package database

import (
	"github.com/Elang89/go-api/config"
	"gopkg.in/mgo.v2"
)

// MongoHostName is a constant for the host name of the mongo service
const MongoHostName string = "mongodb://"

// MongoDbContext contains the database that is used throughout the api
type MongoDbContext struct {
	database *mgo.Session
}

// NewMongoDbContext returns a context object with a database, this is used in repositories
func NewMongoDbContext(c config.Configuration) *MongoDbContext {
	connectionString := generateConnectionString(c)
	db, err := mgo.Dial(connectionString)

	if err != nil {
		panic(err)
	}

	context := MongoDbContext{database: db}
	return &context

}

func generateConnectionString(c config.Configuration) string {
	connectionString := MongoHostName + c.MongoConnection.User + ":" + c.MongoConnection.Pwd + "@" + c.MongoConnection.Host
	return connectionString
}
