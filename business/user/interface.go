package user

import (
	"errors"
)

type Repository interface {
	FindUserByID(id int) (*FindUser, error)
	// We will add Create, Update, Delete here in the next tickets
}

var (
	ErrUserNotFound = errors.New("user not found")
)
