package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//v1 := r.Group("/api")

	r.Run(":1323")
}
