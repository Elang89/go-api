package main

import (
	"github.com/Elang89/go-api/notes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/api")
	notes.RegisterNotes(v1.Group("/notes"))

	r.Run(":5000")
}
