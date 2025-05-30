package repository

import (
	"context"

	"github.com/ortegasixto7/golang-ticket/src/app"
)

type AppRepositoryInterface interface {
	Create(ctx context.Context, app *app.App) error
	FindByID(ctx context.Context, id string) (*app.App, error)
	FindByAppToken(ctx context.Context, appToken string) (*app.App, error)
	Update(ctx context.Context, app *app.App) error
}
