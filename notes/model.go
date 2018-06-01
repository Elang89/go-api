package notes

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

// Note is an object to represent a note saved on the database
type Note struct {
	InternalID bson.ObjectId `bson:"_id" json:"_id"`
	ID         string        `bson:"id" json:"id"`
	Body       string        `bson:"body" json:"body"`
	CreatedOn  time.Time     `bson:"createdOn" json:"createdOn"`
	UpdatedOn  time.Time     `bson:"updatedOn" json:"updatedOn"`
	UserID     string        `bson:"userId" json:"userId"`
}

// NewNote creates a new note struct and returns it
func NewNote(body string, userID string) *Note {
	note := Note{
		ID:        uuid.New().String(),
		Body:      body,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
		UserID:    userID,
	}

	return &note
}
