package notes

import (
	"github.com/Elang89/go-api/config"
	"github.com/Elang89/go-api/database"
	"github.com/globalsign/mgo/bson"
)

var databaseContext = database.NewMongoDbContext(
	config.NewConfiguration())

// GetAllNotes retrieves a collection of which contains all notes from the database
func GetAllNotes() {
	var notes []Note
	err := databaseContext.Notes.Find(bson.M{}).All(&movies)
	return notes, err
}
