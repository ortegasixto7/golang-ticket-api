package repository

import "github.com/ortegasixto7/golang-ticket/src/app"

type AppRepositoryInterface interface {
	Save(app *app.App) (*app.App, error)
	GetByID(id string) (*app.App, error)
	GetByAppToken(appToken string) (*app.App, error)
	Update(app *app.App) error
}
