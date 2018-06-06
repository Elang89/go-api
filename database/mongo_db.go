package database

import (
	"go-api/config"

	"github.com/globalsign/mgo"
)

// MongoDbContext contains the database that is used throughout the api
type MongoDbContext struct {
	Database *mgo.Database
	Notes    *mgo.Collection
}

var db mgo.Database

// NewMongoDbContext returns a context object with a database, this is used in repositories
func NewMongoDbContext(c config.Configuration) *MongoDbContext {
	connectionString := config.GenerateConnectionString(c)
	session, err := mgo.Dial(connectionString)

	if err != nil {
		panic(err)
	}

	context := MongoDbContext{Database: session.DB(c.MongoConnection.Database)}

	context.Notes = getNotesCollection(context)

	return &context

}

// GetNotesCollection returns the notes collection from the database
func getNotesCollection(context MongoDbContext) *mgo.Collection {
	return context.Database.C("Note")
}
