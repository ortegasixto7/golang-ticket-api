package regenerate_token

import (
	"context"
	"errors"
	"time"

	"github.com/ortegasixto7/golang-ticket/src/app/repository"
)

func Execute(ctx context.Context, request *Request, repo repository.AppRepositoryInterface) (*Response, error) {
	integration, err := repo.FindByID(ctx, request.AppID)
	if err != nil {
		return nil, err
	}
	if integration == nil {
		return nil, errors.New("app not found")
	}

	integration.GenerateToken()
	integration.UpdatedAt = time.Now()

	err = repo.Update(ctx, integration)
	if err != nil {
		return nil, err
	}

	return &Response{AppToken: integration.AppToken, ID: integration.ID}, nil
}
