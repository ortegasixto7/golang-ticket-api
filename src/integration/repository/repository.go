package repository

import (
	"github.com/ortegasixto7/golang-ticket/src/database"
	"github.com/ortegasixto7/golang-ticket/src/integration"
)

type IntegrationRepository struct{}

func NewIntegrationRepository() *IntegrationRepository {
	return &IntegrationRepository{}
}

func (repo *IntegrationRepository) Save(integration *integration.Integration) (*integration.Integration, error) {
	modelDB := FromDomain(integration)
	result := database.DB.Create(modelDB)
	if err := result.Error; err != nil {
		return nil, err
	}
	return ToDomain(modelDB), nil
}

func (r *IntegrationRepository) GetByID(id string) (*integration.Integration, error) {
	var modelDB Integration
	result := database.DB.Where("id = ?", id).First(&modelDB)
	if err := result.Error; err != nil {
		return nil, err
	}
	return ToDomain(&modelDB), nil
}

func (r *IntegrationRepository) GetByAppToken(appToken string) (*integration.Integration, error) {
	var modelDB Integration
	result := database.DB.Where("app_token = ?", appToken).First(&modelDB)
	if err := result.Error; err != nil {
		return nil, err
	}
	return ToDomain(&modelDB), nil
}

func (r *IntegrationRepository) Update(integration *integration.Integration) error {
	modelDB := FromDomain(integration)
	result := database.DB.Save(modelDB)
	return result.Error
}
