package notes

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Note is an object to represent a note saved on the database
type Note struct {
	InternalID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Body       string        `bson:"Body" json:"body"`
	CreatedOn  time.Time     `bson:"CreatedOn" json:"createdOn"`
	UpdatedOn  time.Time     `bson:"UpdatedOn" json:"updatedOn"`
	UserID     string        `bson:"UserId" json:"userId"`
}

// NewNote creates a new note struct and returns it
func NewNote() *Note {
	note := Note{
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}

	return &note
}

// CreateNoteForUpdate creates a new note with the data provided in the parameters
func CreateNoteForUpdate() *Note {
	note := Note{
		UpdatedOn: time.Now(),
	}

	return &note
}
