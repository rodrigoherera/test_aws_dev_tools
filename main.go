package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/v1/ping", ping)

	err := router.Run()
	if err != nil {
		// TODO - Add error metric.
		fmt.Println(err)
		panic(err)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
