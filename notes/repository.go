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

// GetNoteByID retrieves a single note from the notes collection from the database
func GetNoteByID(id string) (Note, error) {
	var note Note
	err := databaseContext.Notes.Find(bson.M{"_id": id}).One(&note)
	return note, err
}

// AddNote adds a note to the note collection in the database
func AddNote(note *Note) error {
	newNote := NewNote(note.Body, uuid.New().String())
	err := databaseContext.Notes.Insert(&newNote)
	return err
}

// UpdateNote updates a note in the notes collection in the database
func UpdateNote(note *Note) error {
	err := databaseContext.Notes.Update(bson.M{"_id": note.InternalID}, note)
	return err
}

// RemoveNote removes a note from the notes collection in the database
func RemoveNote(id string) error {
	err := databaseContext.Notes.Remove(bson.M{"_id": id})
	return err
}
