package regenerate_token

import (
	"errors"
	"time"

	"github.com/ortegasixto7/golang-ticket/src/integration/repository"
)

func Handle(request *Request, repo repository.IntegrationRepositoryInterface) (*Response, error) {
	integration, err := repo.GetByID(request.AppID)
	if err != nil {
		return nil, err
	}
	if integration == nil {
		return nil, errors.New("app not found")
	}

	integration.GenerateToken()
	integration.UpdatedAt = time.Now()

	err = repo.Update(integration)
	if err != nil {
		return nil, err
	}

	return &Response{AppToken: integration.AppToken, ID: integration.ID}, nil
}
