package create

import (
	"context"

	"github.com/google/uuid"
	"github.com/ortegasixto7/golang-ticket/src/app"
	"github.com/ortegasixto7/golang-ticket/src/app/repository"
)

func Execute(ctx context.Context, req *Request, repo repository.AppRepositoryInterface) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	integration := &app.App{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
	}
	integration.GenerateToken()
	if err := repo.Create(ctx, integration); err != nil {
		return nil, err
	}

	response := &Response{
		ID:          integration.ID,
		Name:        integration.Name,
		Description: integration.Description,
		AppToken:    integration.AppToken,
	}

	return response, nil
}
