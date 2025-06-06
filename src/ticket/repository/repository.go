package repository

import (
	"github.com/ortegasixto7/golang-ticket/src/database"
	"github.com/ortegasixto7/golang-ticket/src/ticket"
)

type TicketRepository struct{}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (repo *TicketRepository) Save(model *ticket.Ticket) (*ticket.Ticket, error) {
	modelDB := FromDomain(model)
	result := database.DB.Create(modelDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToDomain(modelDB), nil
}

func (repo *TicketRepository) Update(model *ticket.Ticket) (*ticket.Ticket, error) {
	modelDB := FromDomain(model)
	result := database.DB.Save(modelDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToDomain(modelDB), nil
}

func (repo *TicketRepository) GetByID(id int64) (*ticket.Ticket, error) {
	var modelDB Ticket
	result := database.DB.First(&modelDB, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToDomain(&modelDB), nil
}

func (repo *TicketRepository) GetByCode(code string) (*ticket.Ticket, error) {
	var modelDB Ticket
	result := database.DB.Where("code = ?", code).First(&modelDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToDomain(&modelDB), nil
}
