package repository

import (
	"github.com/ortegasixto7/golang-ticket/src/app"
	"github.com/ortegasixto7/golang-ticket/src/database"
)

type AppRepository struct{}

func NewAppRepository() *AppRepository {
	return &AppRepository{}
}

func (repo *AppRepository) Save(integration *app.App) (*app.App, error) {
	modelDB := FromDomain(integration)
	result := database.DB.Create(modelDB)
	if err := result.Error; err != nil {
		return nil, err
	}
	return ToDomain(modelDB), nil
}

func (r *AppRepository) GetByID(id string) (*app.App, error) {
	var modelDB App
	result := database.DB.Where("id = ?", id).First(&modelDB)
	if err := result.Error; err != nil {
		return nil, err
	}
	return ToDomain(&modelDB), nil
}

func (r *AppRepository) GetByAppToken(appToken string) (*app.App, error) {
	var modelDB App
	result := database.DB.Where("app_token = ?", appToken).First(&modelDB)
	if err := result.Error; err != nil {
		return nil, err
	}
	return ToDomain(&modelDB), nil
}

func (r *AppRepository) Update(integration *app.App) error {
	modelDB := FromDomain(integration)
	result := database.DB.Save(modelDB)
	return result.Error
}
