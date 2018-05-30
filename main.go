package main

import (
	"fmt"

	"github.com/Elang89/go-api/config"
	"github.com/Elang89/go-api/database"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//v1 := r.Group("/api")

	config := config.NewConfiguration()
	context := database.NewMongoDbContext(config)

	result := context.Notes.Find()
	fmt.Println(result)

	r.Run(":1323")
}
