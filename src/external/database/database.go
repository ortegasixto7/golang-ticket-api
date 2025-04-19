package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// Asegúrate de que la ruta sea correcta
)

var DB *gorm.DB

func InitializeDatabase() {
	// Es MUY recomendable usar variables de entorno para la información sensible
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Si no están definidas las variables de entorno, usa valores por defecto (solo para desarrollo)
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbUser == "" {
		dbUser = "tu_usuario"
	}
	if dbPassword == "" {
		dbPassword = "tu_contraseña"
	}
	if dbName == "" {
		dbName = "tu_base_de_datos"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos con GORM: %v", err)
	}

	DB = db
	fmt.Println("Conexión a PostgreSQL (GORM) inicializada.")
}
