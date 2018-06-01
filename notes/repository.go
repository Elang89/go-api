package notes

import (
	"github.com/Elang89/go-api/config"
	"github.com/Elang89/go-api/database"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

var databaseContext = database.NewMongoDbContext(
	config.NewConfiguration())

// GetAllNotes retrieves a collection of which contains all notes from the database
func GetAllNotes() ([]Note, error) {
	var notes []Note
	err := databaseContext.Notes.Find(bson.M{}).All(&notes)
	return notes, err
}

// AddNote adds a note to the note collection in the database
func AddNote(note *Note) error {
	newNote := NewNote(note.Body, uuid.New().String())
	err := databaseContext.Notes.Insert(&newNote)
	return err
}
