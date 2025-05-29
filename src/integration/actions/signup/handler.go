package signup

import (
	"github.com/google/uuid"
	"github.com/ortegasixto7/golang-ticket/src/integration"
	"github.com/ortegasixto7/golang-ticket/src/integration/repository"
)

func Handle(req *Request, repo repository.IntegrationRepositoryInterface) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	integration := integration.Integration{
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
