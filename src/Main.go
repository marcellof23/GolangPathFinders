package main

import (
	"net/http"
	"os"

	"github.com/marcellof23/GolangPathFinders/src/services"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
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
	http.ListenAndServe(":"+port, nil)
}
