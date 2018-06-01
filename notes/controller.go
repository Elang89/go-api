package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterNotes registers the routes associated to notes
func RegisterNotes(router *gin.RouterGroup) {
	router.POST("/", Post)
	router.GET("/", GetAll)
	router.PUT("/:id", Put)
	router.GET("/:id", Get)
	router.DELETE("/:id", Delete)
}

// GetAll gets all the records corresponding to notes
func GetAll(ctx *gin.Context) {
	notes, err := GetAllNotes()

	if err != nil {
		panic(err)
	}

	if len(notes) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound,
			"message": "No notes found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
		"data": notes})
}

// Post adds a note record to the database
func Post(ctx *gin.Context) {
	note := NewNote(ctx.PostForm("Body"), ctx.PostForm("UserId"))
	err := AddNote(note)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity,
			"message": "Note could not be inserted", "error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// Put updates a note record in the database
func Put(ctx *gin.Context) {

}

// Get gets one record from the database
func Get(ctx *gin.Context) {

}

// Delete removes a record from the database
func Delete(ctx *gin.Context) {

}
