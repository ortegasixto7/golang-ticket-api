package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ortegasixto7/golang-ticket/src/database"
	integrationSignUp "github.com/ortegasixto7/golang-ticket/src/integration/actions/signup"
	integrationDB "github.com/ortegasixto7/golang-ticket/src/integration/repository"

	"github.com/ortegasixto7/golang-ticket/src/ticket/actions/generate"
	"github.com/ortegasixto7/golang-ticket/src/ticket/actions/validate"

	ticketDB "github.com/ortegasixto7/golang-ticket/src/ticket/repository"
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

	integrationRepo := integrationDB.NewIntegrationRepository()
	integrationSignUpCtrl := integrationSignUp.Controller{
		Repo: integrationRepo,
	}

	ticketRepo := ticketDB.NewTicketRepository()
	ticketGenerateCtrl := generate.Controller{
		Repo: ticketRepo,
	}
	ticketValidateCtrl := validate.Controller{
		Repo: ticketRepo,
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/tickets", ticketGenerateCtrl.Generate)
	router.POST("/tickets/:id/validate", ticketValidateCtrl.Validate)
	router.POST("/apps", integrationSignUpCtrl.SignUp)

	router.Run() // listen and serve on 0.0.0.0:8080
}
