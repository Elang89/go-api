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
func NewNote(body string, userID string) *Note {
	note := Note{
		Body:      body,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
		UserID:    userID,
	}

	return &note
}

// MakeNoteFromData creates a new note with the data provided in the parameters
func MakeNoteFromData(internalID string, body string, createdOn string, userID string) *Note {
	layout := "2006-01-02T15:04:05-0700"
	date, err := time.Parse(layout, createdOn)

	if err != nil {
		panic(err)
	}

	note := Note{
		InternalID: bson.ObjectIdHex(internalID),
		Body:       body,
		CreatedOn:  date,
		UpdatedOn:  time.Now(),
		UserID:     userID,
	}

	return &note
}
