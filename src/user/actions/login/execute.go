package login

import (
	"context"
	"fmt"

	"github.com/ortegasixto7/golang-ticket/src/auth"
	"github.com/ortegasixto7/golang-ticket/src/user"
	"github.com/ortegasixto7/golang-ticket/src/user/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = fmt.Errorf("invalid credentials")

func Execute(ctx context.Context, req *Request, repo repository.UserRepositoryInterface) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	foundUser, err := repo.FindByUsername(ctx, req.Username)
	if err != nil {
		if err == user.ErrUserNotFound {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := auth.GenerateToken(foundUser.ID, foundUser.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &Response{
		Token: token,
	}, nil
}
