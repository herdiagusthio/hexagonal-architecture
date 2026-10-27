package user

import (
	"errors"
)

// Service is the inbound port: what the API layer may ask of the user domain.
type Service interface {
	FindUserByID(id int) (*FindUser, error)
}

// Repository is the outbound port: what the domain needs from storage, not how it is stored.
type Repository interface {
	FindUserByID(id int) (*FindUser, error)
	// We will add Create, Update, Delete here in the next tickets
}

var (
	ErrUserNotFound = errors.New("user not found")
)
