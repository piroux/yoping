package models

import (
	"testing"
)

// CAUTION: Copilot generated tests

func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		name          string
		phoneNumber   string
		expectedError bool
	}{
		{
			name:          "Valid French phone number",
			phoneNumber:   "+33612345678", // Valid French mobile number
			expectedError: false,
		},
		{
			name:          "Invalid phone number",
			phoneNumber:   "12345", // Clearly invalid
			expectedError: true,
		},
		{
			name:          "Empty phone number",
			phoneNumber:   "",
			expectedError: true,
		},
		{
			name:          "Valid US phone number",
			phoneNumber:   "+14155552671", // Valid US number
			expectedError: false,
		},
		{
			name:          "Invalid format phone number",
			phoneNumber:   "+33-6-12-34-56-78", // Valid number but invalid format for parsing
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePhoneNumber(tt.phoneNumber)
			if (err != nil) != tt.expectedError {
				t.Errorf("validatePhoneNumber(%q) error = %v, expectedError = %v", tt.phoneNumber, err, tt.expectedError)
			}
		})
	}
}
