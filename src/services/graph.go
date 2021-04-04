package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _ := c.FormFile("inpFile")
		log.Println(file.Filename)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, map[string]string{
			"data": "hello from graph",
		})
	}
}
