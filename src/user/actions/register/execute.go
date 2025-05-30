package register

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ortegasixto7/golang-ticket/src/user"
	"github.com/ortegasixto7/golang-ticket/src/user/repository"
	"golang.org/x/crypto/bcrypt"
)

func Execute(ctx context.Context, req *Request, repo repository.UserRepositoryInterface) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	existingUser, err := repo.FindByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, user.ErrUserNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, user.ErrUsernameTaken
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	newUser := &user.User{
		ID:        uuid.New(),
		Name:      req.Name,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := repo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return &Response{
		ID:       newUser.ID,
		Name:     newUser.Name,
		Username: newUser.Username,
	}, nil
}
