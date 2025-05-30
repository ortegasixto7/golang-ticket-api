package repository

import (
	"context"
	"fmt"

	"github.com/ortegasixto7/golang-ticket/src/app"
	"github.com/ortegasixto7/golang-ticket/src/database"
	"gorm.io/gorm"
)

type appRepository struct {
	db *gorm.DB
}

func NewAppRepository() AppRepositoryInterface {
	return &appRepository{
		db: database.DB,
	}
}

func (r *appRepository) Create(ctx context.Context, app *app.App) error {
	if err := r.db.WithContext(ctx).Create(app).Error; err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}
	return nil
}

func (r *appRepository) FindByID(ctx context.Context, id string) (*app.App, error) {
	var foundApp app.App
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&foundApp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app.ErrAppNotFound
		}
		return nil, fmt.Errorf("failed to find app by id: %w", err)
	}
	return &foundApp, nil
}

func (r *appRepository) FindByAppToken(ctx context.Context, appToken string) (*app.App, error) {
	var foundApp app.App
	err := r.db.WithContext(ctx).Where("app_token = ? AND deleted_at IS NULL", appToken).First(&foundApp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app.ErrAppNotFound
		}
		return nil, fmt.Errorf("failed to find app by token: %w", err)
	}
	return &foundApp, nil
}

func (r *appRepository) Update(ctx context.Context, appToUpdate *app.App) error {
	result := r.db.WithContext(ctx).Save(appToUpdate)
	if result.Error != nil {
		return fmt.Errorf("failed to update app: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return app.ErrAppNotFound
	}
	return nil
}
