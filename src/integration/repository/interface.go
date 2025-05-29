package repository

import "github.com/ortegasixto7/golang-ticket/src/integration"

type IntegrationRepositoryInterface interface {
	Save(integration *integration.Integration) (*integration.Integration, error)
	GetByID(id string) (*integration.Integration, error)
	GetByAppToken(appToken string) (*integration.Integration, error)
	Update(integration *integration.Integration) error
}
