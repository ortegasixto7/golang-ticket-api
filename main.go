package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	appCreate "github.com/ortegasixto7/golang-ticket/src/app/actions/create"
	appRegenToken "github.com/ortegasixto7/golang-ticket/src/app/actions/regenerate_token"
	"github.com/ortegasixto7/golang-ticket/src/database"

	appDB "github.com/ortegasixto7/golang-ticket/src/app/repository"
	userLogin "github.com/ortegasixto7/golang-ticket/src/user/actions/login"
	userRegister "github.com/ortegasixto7/golang-ticket/src/user/actions/register"

	userDB "github.com/ortegasixto7/golang-ticket/src/user/repository"

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

	appRepo := appDB.NewAppRepository()
	appCreateCtrl := appCreate.Controller{
		Repo: appRepo,
	}
	appRegenTokenCtrl := appRegenToken.Controller{
		Repo: appRepo,
	}

	ticketRepo := ticketDB.NewTicketRepository()
	ticketGenerateCtrl := generate.Controller{
		Repo: ticketRepo,
	}
	ticketValidateCtrl := validate.Controller{
		Repo: ticketRepo,
	}

	userRepo := userDB.NewUserRepository()
	userRegisterCtrl := userRegister.NewController(userRepo)
	userLoginCtrl := userLogin.NewController(userRepo)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/tickets", ticketGenerateCtrl.Generate)
	router.POST("/tickets/validate", ticketValidateCtrl.Validate)
	router.POST("/apps", appCreateCtrl.Handle)
	router.POST("/apps/tokens/regenerate", appRegenTokenCtrl.Handle)
	router.POST("/users/register", userRegisterCtrl.Handle)
	router.POST("/users/login", userLoginCtrl.Handle)

	router.Run() // listen and serve on 0.0.0.0:8080
}
