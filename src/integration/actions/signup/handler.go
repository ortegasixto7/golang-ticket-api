package signup

import (
	"time"

	"github.com/ortegasixto7/golang-ticket/src/integration"
	"github.com/ortegasixto7/golang-ticket/src/integration/database"
)

func Execute(req *Request, repo *database.IntegrationRepository) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	integration := integration.Integration{
		ID:          time.Now().Format(time.RFC3339),
		Name:        req.Name,
		Description: req.Description,
	}
	integration.GenerateToken()
	if _, err := repo.Insert(&integration); err != nil {
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
