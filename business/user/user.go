package user

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidUsername = errors.New("username must be 3-20 alphanumeric characters")
)

// Email Value Object
type Email string

func NewEmail(e string) (Email, error) {
	e = strings.TrimSpace(e)
	if e == "" {
		return "", ErrInvalidEmail
	}
	// Simple email regex
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !re.MatchString(strings.ToLower(e)) {
		return "", ErrInvalidEmail
	}
	return Email(e), nil
}

func (e Email) String() string {
	return string(e)
}

// Username Value Object
type Username string

func NewUsername(u string) (Username, error) {
	u = strings.TrimSpace(u)
	if u == "" {
		return "", ErrInvalidUsername
	}
	// 3-20 alphanumeric characters
	re := regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
	if !re.MatchString(u) {
		return "", ErrInvalidUsername
	}
	return Username(u), nil
}

func (u Username) String() string {
	return string(u)
}

type User struct {
	ID          int
	Name        string
	Email       Email
	PhoneNumber string
	Username    Username
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type FindUser struct {
	ID          int
	Name        string
	Email       Email
	PhoneNumber string
	Username    Username
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// NewUser create new User with validated value objects
func NewUser(
	id int,
	name string,
	emailStr string,
	phoneNumber string,
	usernameStr string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {
	email, err := NewEmail(emailStr)
	if err != nil {
		return nil, fmt.Errorf("email validation: %w", err)
	}

	username, err := NewUsername(usernameStr)
	if err != nil {
		return nil, fmt.Errorf("username validation: %w", err)
	}

	return &User{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Username:    username,
		Password:    password,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   nil,
	}, nil
}

// ModifyUser update existing User data using value objects for validation
func (oldData *FindUser) ModifyUser(newName string, newPhoneNumber string, modifiedAt time.Time) (*FindUser, error) {
	name := newName
	if name == "" {
		name = oldData.Name
	}

	phone := newPhoneNumber
	if phone == "" {
		phone = oldData.PhoneNumber
	}

	return &FindUser{
		ID:          oldData.ID,
		Name:        name,
		Email:       oldData.Email,
		PhoneNumber: phone,
		Username:    oldData.Username,
		CreatedAt:   oldData.CreatedAt,
		UpdatedAt:   modifiedAt,
		DeletedAt:   nil,
	}, nil
}
