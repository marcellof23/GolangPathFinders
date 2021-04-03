package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Wakwaw() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"data": "marcello wakwaw",
		})
	}
}
