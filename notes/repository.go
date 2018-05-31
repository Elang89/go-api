package notes

import (
	"github.com/Elang89/go-api/config"
	"github.com/Elang89/go-api/database"
)

var databaseContext = database.NewMongoDbContext(
	config.NewConfiguration())

// GetAllNotes retrieves a collection of which contains all notes from the database
func GetAllNotes() {
	var notes []Note
}
