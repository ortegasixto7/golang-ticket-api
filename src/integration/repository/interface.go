package repository

import "github.com/ortegasixto7/golang-ticket/src/integration"

type IntegrationRepositoryInterface interface {
	Save(integration *integration.Integration) (*integration.Integration, error)
	GetByID(id int64) (*integration.Integration, error)
}
