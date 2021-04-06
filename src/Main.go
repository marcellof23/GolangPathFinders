package main

import (
	"AStarPathFinder/src/services"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/wakwaw", services.Wakwaw())
	api.GET("/test", services.Test())
	api.GET("/calc", services.Calc())
	api.POST("/graph", services.ParseFile())
	api.POST("/graphdata", services.PostData())
	router.Run("127.0.0.1:5000")
}
