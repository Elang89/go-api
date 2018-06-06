package notes

import (
	"go-api/config"
	"go-api/database"

	"github.com/globalsign/mgo/bson"
)

var databaseContext = database.NewMongoDbContext(
	config.NewConfiguration())

// GetAllNotes retrieves a collection of which contains all notes from the database
func GetAllNotes() ([]Note, error) {
	var notes []Note
	err := databaseContext.Notes.Find(bson.M{}).All(&notes)
	return notes, err
}

// GetNoteByID retrieves a single note from the notes collection from the database
func GetNoteByID(id string) (Note, error) {
	var note Note
	err := databaseContext.Notes.FindId(bson.ObjectIdHex(id)).One(&note)
	return note, err
}

// AddNote adds a note to the note collection in the database
func AddNote(note *Note) error {
	err := databaseContext.Notes.Insert(note)
	return err
}

// UpdateNote updates a note in the notes collection in the database
func UpdateNote(id string, note *Note) error {
	err := databaseContext.Notes.UpdateId(bson.ObjectIdHex(id), note)
	return err
}

// RemoveNote removes a note from the notes collection in the database
func RemoveNote(id string) error {
	err := databaseContext.Notes.RemoveId(bson.ObjectIdHex(id))
	return err
}
