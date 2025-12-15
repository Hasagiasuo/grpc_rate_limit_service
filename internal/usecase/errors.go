package usecase

import "errors"

var (
	ErrCannotUpdateService = errors.New("cannot update service")
	ErrInsufficientFunds = errors.New("insufficient funds")
)
