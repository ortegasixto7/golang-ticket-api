package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ortegasixto7/golang-ticket/src/controllers"
)

func main() {

	var ticketCtrl controllers.TicketController

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/tickets", ticketCtrl.Health)

	router.Run() // listen and serve on 0.0.0.0:8080
}
