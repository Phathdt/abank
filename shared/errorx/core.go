package errorx

import "errors"

var (
	ErrCannotListAccounts = errors.New("cannot list accounts")
	ErrGetUser            = errors.New("cannot get user")
	ErrGetAccount         = errors.New("cannot get account")
)
