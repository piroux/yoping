package models

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/nyaruka/phonenumbers"
)

const phoneNumberDefaultRegion = "FR"

type PhoneNumberPair struct {
	From, To string // PK
}

func NewPhoneNumberPair(fromPhoneNumber, toPhoneNumber string) (*PhoneNumberPair, error) {
	var err error

	m := &PhoneNumberPair{
		From: fromPhoneNumber,
		To:   toPhoneNumber,
	}

	err = m.validate()
	if err != nil {
		return nil, fmt.Errorf("failed to create PhoneNumberPair: %w", err)
	}

	return m, nil
}

func (m PhoneNumberPair) validate() error {
	var (
		result error
		err    error
	)

	err = validatePhoneNumber(m.From)
	if err != nil {
		result = multierror.Append(result, fmt.Errorf("failed to validate From phone number: %w", err))
	}
	err = validatePhoneNumber(m.To)
	if err != nil {
		result = multierror.Append(result, fmt.Errorf("failed to validate To phone number: %w", err))
	}

	return result
}

func validatePhoneNumber(phoneNumber string) error {
	pn, err := phonenumbers.Parse(phoneNumber, phoneNumberDefaultRegion)
	if err != nil {
		return err
	}

	ok := phonenumbers.IsValidNumber(pn)
	if !ok {
		return errors.New("invalid phone number")
	}

	return nil
}
