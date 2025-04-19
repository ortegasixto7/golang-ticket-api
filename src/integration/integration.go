package integration

import (
	"strings"

	"github.com/google/uuid"
)

type Integration struct {
	ID          string
	Name        string
	Description string
	AppToken    string
	IsEnabled   bool
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
}

func (i *Integration) GenerateToken() string {
	token := uuid.New()
	tokenString := token.String()
	tokenString = strings.ReplaceAll(tokenString, "-", "")
	i.AppToken = strings.ToUpper(tokenString)
	return i.AppToken
}
