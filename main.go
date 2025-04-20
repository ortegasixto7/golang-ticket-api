package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ortegasixto7/golang-ticket/src/controllers"
	"github.com/ortegasixto7/golang-ticket/src/database"
	integrationSignUp "github.com/ortegasixto7/golang-ticket/src/integration/actions/signup"
	integrationDB "github.com/ortegasixto7/golang-ticket/src/integration/repository"
)

func main() {

	// Carga las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		// Si no se encuentra el archivo .env, la aplicación puede seguir funcionando
		// utilizando las variables de entorno del sistema.
	}

	fmt.Println("Starting app ...")

	database.InitDatabase()

	var ticketCtrl controllers.TicketController

	var integrationRepo integrationDB.IntegrationRepositoryInterface
	integrationRepo = integrationDB.NewIntegrationRepository()
	integrationSignUpCtrl := integrationSignUp.Controller{
		Repo: &integrationRepo,
	}

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
