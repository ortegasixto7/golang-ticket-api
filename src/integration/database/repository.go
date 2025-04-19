package database

import (
	"github.com/ortegasixto7/golang-ticket/src/external/database"
	"github.com/ortegasixto7/golang-ticket/src/integration"
)

type IntegrationRepository struct{}

func NewIntegrationRepository() *IntegrationRepository {
	return &IntegrationRepository{}
}

func (repo *IntegrationRepository) Insert(integration *integration.Integration) (*integration.Integration, error) {
	modelDB := FromDomain(integration)
	result := database.DB.Create(modelDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToDomain(modelDB), nil
}

func (r *IntegrationRepository) GetByID(id int64) (*integration.Integration, error) {
	var modelDB Integration
	result := database.DB.First(&modelDB, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToDomain(&modelDB), nil
}
