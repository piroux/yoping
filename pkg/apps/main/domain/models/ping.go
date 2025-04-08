package models

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-multierror"
)

type Ping struct {
	PhoneNumbers PhoneNumberPair // PK
	TimeCreated  time.Time

	// TODO. Upgrade to
	//TimeRxed  time.Time
	//TimeTxed  time.Time
}

type PingKey = PhoneNumberPair

type PingData struct {
	PhoneNumberFrom, PhoneNumberTo string
}

func NewPing(fromPhoneNumber, toPhoneNumber string) (*Ping, error) {
	var err error

	now := time.Now().UTC()

	pnPair, err := NewPhoneNumberPair(fromPhoneNumber, toPhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to create Ping: failed to create PhoneNumberPair for Ping:%w", err)
	}

	m := &Ping{
		PhoneNumbers: *pnPair,
		TimeCreated:  now,
	}

	err = m.validate()
	if err != nil {
		return nil, fmt.Errorf("failed to create Ping: failed to validate Ping: %w", err)
	}

	return m, nil
}

func NewPingWithPhoneNumberPair(phoneNumberPair PhoneNumberPair) (*Ping, error) {
	var err error

	now := time.Now().UTC()

	m := &Ping{
		PhoneNumbers: phoneNumberPair,
		TimeCreated:  now,
	}

	err = m.validate()
	if err != nil {
		return nil, fmt.Errorf("failed to create Ping: failed to validate Ping: %w", err)
	}

	return m, err
}

func (m Ping) validate() error {
	var (
		result error
	)

	okTimeCreated := !m.TimeCreated.IsZero()
	if !okTimeCreated {
		result = multierror.Append(result, fmt.Errorf("TimeCreated is invalid"))
	}

	return result
}
