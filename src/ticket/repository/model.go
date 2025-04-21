package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ortegasixto7/golang-ticket/src/ticket"
)

type Ticket struct {
	ID        string     `gorm:"type:varchar(100);not null"`
	Title     string     `gorm:"type:varchar(100);not null"`
	Code      string     `gorm:"type:varchar(100);not null;unique"`
	CreatedAt time.Time  `gorm:"type:timestamptz;not null"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;not null"`
	DeletedAt *time.Time `gorm:"type:timestamptz"`
	ExpiresAt *time.Time `gorm:"type:timestamptz"`
}

func FromDomain(model *ticket.Ticket) *Ticket {
	return &Ticket{
		ID:        model.ID.String(),
		Title:     model.Title,
		Code:      model.Code,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
		ExpiresAt: model.ExpiresAt,
	}
}

func ToDomain(model *Ticket) *ticket.Ticket {
	idUUID, err := uuid.Parse(model.ID)
	if err != nil {
		fmt.Println("Error al convertir el string a UUID:", err)
		return nil
	}
	return &ticket.Ticket{
		ID:        idUUID,
		Title:     model.Title,
		Code:      model.Code,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
		ExpiresAt: model.ExpiresAt,
	}
}
