package app

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type App struct {
	ID          string
	Name        string
	Description string
	AppToken    string
	IsEnabled   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (i *App) GenerateToken() string {
	token := uuid.New()
	tokenString := token.String()
	tokenString = strings.ReplaceAll(tokenString, "-", "")
	i.AppToken = strings.ToUpper(tokenString)
	return i.AppToken
}
