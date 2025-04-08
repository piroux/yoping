package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type User struct {
	Id          uuid.UUID // PK
	NameFull    string
	PhoneNumber string
}

type UserKey = uuid.UUID

type UserData struct {
	NameFull    string
	PhoneNumber string
}

func NewUser(name string, phoneNumber string) (*User, error) {
	var err error

	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to create User: failed to create uuid for User: %w", err)
	}
	m := &User{
		Id:          id,
		NameFull:    name,
		PhoneNumber: phoneNumber,
	}

	err = m.validate()
	if err != nil {
		return nil, fmt.Errorf("failed to create User: failed to validate User: %w", err)
	}

	return m, err
}

func (m User) validate() error {
	var (
		result error
		err    error
	)

	err = validatePhoneNumber(m.PhoneNumber)
	if err != nil {
		result = multierror.Append(result, err)
	}

	if len(m.NameFull) == 0 {
		result = multierror.Append(result, errors.New("failed to validate Name: empty"))
	}

	return result
}
