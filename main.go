package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ortegasixto7/golang-ticket/src/controllers"
	integrationSignUp "github.com/ortegasixto7/golang-ticket/src/integration/actions/signup"
)

func main() {

	var ticketCtrl controllers.TicketController
	var integrationSignUpCtrl integrationSignUp.Controller

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/tickets", ticketCtrl.Health)
	router.POST("/tickets", ticketCtrl.Generate)
	router.POST("/apps", integrationSignUpCtrl.SignUp)

	router.Run() // listen and serve on 0.0.0.0:8080
}
