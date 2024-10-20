package user

import "errors"

var (
	ErrGetUser               = errors.New("store get user")
	ErrInsertUserIfNotExists = errors.New("store insert user if not exists")
)
