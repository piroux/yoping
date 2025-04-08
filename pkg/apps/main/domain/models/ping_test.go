package models

import (
	"testing"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

// CAUTION: Copilot generated tests

func TestPing_Validate(t *testing.T) {
	t.Run("valid Ping", func(t *testing.T) {
		ping := Ping{
			TimeCreated: time.Now().UTC(),
		}

		err := ping.validate()
		assert.NoError(t, err, "expected no error for valid Ping")
	})

	t.Run("invalid Ping with zero TimeCreated", func(t *testing.T) {
		ping := Ping{
			TimeCreated: time.Time{}, // Zero value
		}

		err := ping.validate()
		assert.Error(t, err, "expected an error for invalid Ping with zero TimeCreated")

		multiErr, ok := err.(*multierror.Error)
		assert.True(t, ok, "expected error to be of type *multierror.Error")
		assert.Contains(t, multiErr.Error(), "TimeCreated is invalid", "expected error message to contain 'TimeCreated is invalid'")
	})
}

func TestNewPing(t *testing.T) {
	t.Run("valid Ping creation", func(t *testing.T) {
		fromPhoneNumber := "+33600000001"
		toPhoneNumber := "+33600000003"

		ping, err := NewPing(fromPhoneNumber, toPhoneNumber)

		assert.NoError(t, err, "expected no error for valid Ping creation")
		assert.NotNil(t, ping, "expected Ping to be created")
		assert.Equal(t, fromPhoneNumber, ping.PhoneNumbers.From, "expected From phone number to match")
		assert.Equal(t, toPhoneNumber, ping.PhoneNumbers.To, "expected To phone number to match")
		assert.False(t, ping.TimeCreated.IsZero(), "expected TimeCreated to be set")
	})

	t.Run("invalid Ping creation with invalid phone numbers", func(t *testing.T) {
		fromPhoneNumber := "invalid"
		toPhoneNumber := "+0987654321"

		ping, err := NewPing(fromPhoneNumber, toPhoneNumber)

		assert.Error(t, err, "expected an error for invalid phone numbers")
		assert.Nil(t, ping, "expected Ping to be nil on error")
		assert.Contains(t, err.Error(), "failed to create PhoneNumberPair", "expected error message to contain 'failed to create PhoneNumberPair'")
	})

	t.Run("invalid Ping creation with validation failure", func(t *testing.T) {
		// Simulate a validation failure by mocking NewPhoneNumberPair to return a valid pair
		// and then manually setting TimeCreated to zero.
		fromPhoneNumber := "+33600000001"
		toPhoneNumber := "+33600000003"

		ping, err := NewPing(fromPhoneNumber, toPhoneNumber)
		assert.NoError(t, err, "expected no error for valid phone numbers")
		assert.NotNil(t, ping, "expected Ping to be created")

		// Manually set TimeCreated to zero to simulate validation failure
		ping.TimeCreated = time.Time{}
		err = ping.validate()

		assert.Error(t, err, "expected an error for validation failure")
		assert.Contains(t, err.Error(), "TimeCreated is invalid", "expected error message to contain 'TimeCreated is invalid'")
	})
}
