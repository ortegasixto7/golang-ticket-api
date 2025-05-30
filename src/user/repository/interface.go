package repository

import (
	"context"

	"github.com/ortegasixto7/golang-ticket/src/user"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *user.User) error
	FindByUsername(ctx context.Context, username string) (*user.User, error)
}
