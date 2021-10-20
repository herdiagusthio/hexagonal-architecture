package user

import "time"

type User struct {
	ID          int
	Name        string
	Email       string
	PhoneNumber string
	Username    string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type FindUser struct {
	ID          int
	Name        string
	Email       string
	PhoneNumber string
	Username    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//NewUser create new User
func NewUser(
	id int,
	name string,
	email string,
	phoneNumber string,
	username string,
	password string,
	createdAt time.Time,
	updatedAt time.Time) User {

	return User{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Username:    username,
		Password:    password,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   nil,
	}
}

//ModifyUser update existing User data
func (oldData *FindUser) ModifyUser(newName string, newPhoneNumber string, modifiedAt time.Time, updater string) FindUser {
	var name, phone string
	name = newName
	phone = newPhoneNumber

	if newName == "" {
		name = oldData.Name
	}

	if newPhoneNumber == "" {
		phone = oldData.PhoneNumber
	}

	return FindUser{
		ID:          oldData.ID,
		Name:        name,
		Email:       oldData.Email,
		PhoneNumber: phone,
		Username:    oldData.Username,
		CreatedAt:   oldData.CreatedAt,
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
}
