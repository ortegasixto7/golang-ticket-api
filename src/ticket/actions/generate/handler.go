package generate

import (
	"github.com/google/uuid"
	"github.com/ortegasixto7/golang-ticket/src/ticket"
	"github.com/ortegasixto7/golang-ticket/src/ticket/repository"
)

func Execute(req *Request, repo repository.TicketRepositoryInterface) (*Response, error) {
	tkt := ticket.Ticket{
		ID:        uuid.New(),
		Title:     req.Title,
		ExpiresAt: req.ExpiresAt,
	}
	tkt.GenerateCode()

	if _, err := repo.Save(&tkt); err != nil {
		return nil, err
	}

	response := &Response{
		ID:        tkt.ID.String(),
		Title:     tkt.Title,
		Code:      tkt.Code,
		ExpiresAt: tkt.ExpiresAt,
	}

	return response, nil
}
