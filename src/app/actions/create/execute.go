package create

import (
	"github.com/google/uuid"
	"github.com/ortegasixto7/golang-ticket/src/app"
	"github.com/ortegasixto7/golang-ticket/src/app/repository"
)

func Execute(req *Request, repo repository.AppRepositoryInterface) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	integration := app.App{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
	}
	integration.GenerateToken()
	if _, err := repo.Save(&integration); err != nil {
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
