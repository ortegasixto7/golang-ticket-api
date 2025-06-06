package register

import (
	"github.com/google/uuid"
)

type Response struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
}
