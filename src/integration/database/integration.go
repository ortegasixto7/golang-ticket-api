package database

import (
	"github.com/ortegasixto7/golang-ticket/src/integration"
)

type Integration struct {
	ID          string `gorm:"type:varchar(100);not null"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	AppToken    string `gorm:"type:varchar(100);not null;unique"`
	IsEnabled   bool   `gorm:"type:boolean;default:true"`
	CreatedAt   int64  `gorm:"type:bigint;not null"`
	UpdatedAt   int64  `gorm:"type:bigint;not null"`
	DeletedAt   int64  `gorm:"type:bigint;not null"`
}

func FromDomain(i *integration.Integration) *Integration {
	return &Integration{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		AppToken:    i.AppToken,
		IsEnabled:   i.IsEnabled,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		DeletedAt:   i.DeletedAt,
	}
}

func ToDomain(model *Integration) *integration.Integration {
	return &integration.Integration{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		AppToken:    model.AppToken,
		IsEnabled:   model.IsEnabled,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt,
	}
}
