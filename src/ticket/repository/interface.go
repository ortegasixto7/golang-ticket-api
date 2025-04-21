package repository

import (
	"github.com/ortegasixto7/golang-ticket/src/ticket"
)

type TicketRepositoryInterface interface {
	Save(model *ticket.Ticket) (*ticket.Ticket, error)
	GetByID(id int64) (*ticket.Ticket, error)
}
