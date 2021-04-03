package main

import (
	"AStarPathFinder/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/wakwaw", services.Wakwaw())
	r.GET("/test", services.Wakwaw())
	r.Run("127.0.0.1:5000")
}
