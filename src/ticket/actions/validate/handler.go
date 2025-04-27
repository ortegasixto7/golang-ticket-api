package validate

import (
	"errors"
	"time"

	"github.com/ortegasixto7/golang-ticket/src/ticket/repository"
)

func Execute(req *Request, repo repository.TicketRepositoryInterface) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	tkt, tkterr := repo.GetByCode(req.Code)
	if tkterr != nil {
		return nil, tkterr
	}

	res := &Response{
		ID:      tkt.ID.String(),
		Title:   tkt.Title,
		Code:    tkt.Code,
		IsValid: tkt.IsExpired(),
	}

	if tkt.ExpiresAt.Before(time.Now()) {
		return res, errors.New("ticket expired")
	}

	now := time.Now()
	tkt.ExpiresAt = &now

	if _, err := repo.Update(tkt); err != nil {
		return nil, err
	}

	return res, nil
}
