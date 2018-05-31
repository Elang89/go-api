package main

import (
	"context"

	"github.com/Elang89/go-api/config"
	"github.com/Elang89/go-api/database"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//v1 := r.Group("/api")

	config := config.NewConfiguration()
	dataContext := database.NewMongoDbContext(config)

	result := dataContext.Notes.Find(context.Background())

	r.Run(":1323")
}
