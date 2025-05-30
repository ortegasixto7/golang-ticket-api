package repository

import (
	"context"

	"github.com/ortegasixto7/golang-ticket/src/database"
	"github.com/ortegasixto7/golang-ticket/src/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepositoryInterface {
	return &userRepository{
		db: database.DB,
	}
}

func (r *userRepository) Create(ctx context.Context, user *user.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	var foundUser user.User
	err := r.db.WithContext(ctx).Where("username = ? AND deleted_at IS NULL", username).First(&foundUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return &foundUser, nil
}
