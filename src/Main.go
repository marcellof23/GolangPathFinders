package main

import (
	"AStarPathFinder/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/wakwaw", services.Wakwaw())
	r.GET("/test", services.Test())
	r.GET("/calc", services.Calc())
	r.Run("127.0.0.1:5000")
}
