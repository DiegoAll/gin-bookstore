package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// RESTful API usually give the response in JSON format, but there are several types of response to the client.
		c.JSON(http.StatusOK, gin.H{"data": "Que mas pues"})
	})

	r.Run()
}
