package app

import "errors"

var (
	ErrAppNotFound   = errors.New("app not found")
	ErrAppTokenTaken = errors.New("app token already taken")
	ErrInvalidInput  = errors.New("invalid input")
)
