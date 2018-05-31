package notes

import (
	"time"

	"github.com/google/uuid"
)

// Note is an object to represent a note saved on the database
type Note struct {
	InternalID string    `json:"_id"`
	ID         string    `json:"id"`
	Body       string    `json:"body"`
	CreatedOn  time.Time `json:"createdOn"`
	UpdatedOn  time.Time `json:"updatedOn"`
	UserID     string    `json:"userId"`
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
