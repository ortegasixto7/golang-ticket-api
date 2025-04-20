package database

import (
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	databaseURL := os.Getenv("DB_URL")
	if databaseURL == "" {
		log.Fatal("La variable de entorno DB_URL no está configurada.")
		return
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos con GORM: %v", err)
	}

	DB = db
	fmt.Println("Conexión a PostgreSQL (GORM) inicializada.")

	// Obtener el *sql.DB subyacente de la conexión GORM
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error al obtener la conexión *sql.DB de GORM: %v", err)
		return
	}

	// Configurar el dialecto para goose
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set dialect: %v\n", err)
		return
	}

	// Pasar el *sql.DB a goose.Up
	if err := goose.Up(sqlDB, "./migrations"); err != nil {
		log.Fatalf("failed to run migrations: %v\n", err)
	}

	fmt.Println("Database migrations applied successfully!")
}
