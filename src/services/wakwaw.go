package services

import (
	"AStarPathFinder/src/Astar"
	"AStarPathFinder/src/Models"
	"AStarPathFinder/src/algorithm"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Wakwaw() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, map[string]string{
			"data": "marcello wakwaw",
		})
	}
}

func Test() gin.HandlerFunc {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return func(c *gin.Context) {
		file, err := os.Open("../constants/peta.md")
		if err != nil {
			fmt.Println("Path" + dir)
			fmt.Println(err)
		}
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		b, err := ioutil.ReadAll(file)
		fmt.Print(b)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, map[string]string{
			"data": "jesson wakwaw",
		})
	}
}

func Calc() gin.HandlerFunc {
	algorithm.TestDfs()
	Astar.PrintAstar()
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, map[string]string{
			"data": "marcello wakwaws",
			// "jesson": Astar.StringAstar(),
		})
		c.JSON(http.StatusOK, map[string]string{
			"datas": "hello from yey",
		})
	}
}

func PostData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var astardata Models.AstarData
		c.BindJSON(&astardata)
		if Astar.StringAstars(astardata) == "" {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, map[string]string{
				"Path":      Astar.StringAstars(astardata),
				"Distance":  Astar.HeuristicHaversineJson(),
				"Bobotedge": Astar.EdgeJson(),
			})
		}
	}
}
