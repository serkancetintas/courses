package domain

import (
	"errors"
	"time"
)

type User struct {
	id          string
	firstName   string
	lastName    string
	email       string
	createdTime time.Time
}

func NewUser(id string, firstName string, lastName string, email string, createdTime time.Time) (*User, error) {
	if id == "" {
		return nil, errors.New("empty user id")
	}
	if firstName == "" {
		return nil, errors.New("empty first name")
	}
	if lastName == "" {
		return nil, errors.New("empty last name")
	}
	if email == "" {
		return nil, errors.New("empty email")
	}
	if createdTime.IsZero() {
		return nil, errors.New("zero created time")
	}

	return &User{
		id:          id,
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		createdTime: createdTime,
	}, nil
}
