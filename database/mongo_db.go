package database

import (
	"github.com/Elang89/go-api/config"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoDbContext contains the database that is used throughout the api
type MongoDbContext struct {
	Database *mongo.Database
	Notes    *mongo.Collection
}

var db mongo.Database

// NewMongoDbContext returns a context object with a database, this is used in repositories
func NewMongoDbContext(c config.Configuration) MongoDbContext {
	connectionString := GenerateConnectionString(c)
	db, err := mongo.NewClient(connectionString)

	if err != nil {
		panic(err)
	}

	context := MongoDbContext{Database: db.Database(c.MongoConnection.Database)}

	context.Notes = getNotesCollection(context)

	return &context

}

// GetNotesCollection returns the notes collection from the database
func getNotesCollection(context MongoDbContext) *mongo.Collection {
	return db.Collection("Note")
}
