package notes

import (
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

}

// Post adds a note record to the database
func Post(ctx *gin.Context) {

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
