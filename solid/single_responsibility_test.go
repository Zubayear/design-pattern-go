// Package solid demonstrates the Single Responsibility Principle (SRP)
package solid_test

import (
	"errors"
	"testing"

	"github.com/Zubayear/design-pattern-golang/solid"
)

type MockValidator struct{}

func (mv *MockValidator) Validate(user *solid.User) error {
	if user.Email == "" || user.FirstName == "" || user.LastName == "" {
		return errors.New("invalid user")
	}
	return nil
}

func TestUserValidator_ValidUser(t *testing.T) {
	validator := MockValidator{}
	user := &solid.User{
		ID:        1,
		FirstName: "Alice",
		LastName:  "Smith",
		Email:     "alice@example.com",
	}

	if err := validator.Validate(user); err != nil {
		t.Errorf("Expected user to be valid, go error: %v", err)
	}
}

func TestUserValidator_InvalidUser(t *testing.T) {
	validator := MockValidator{}
	user := &solid.User{
		ID:        2,
		FirstName: "",
		LastName:  "Smith",
		Email:     "",
	}

	err := validator.Validate(user)
	if err == nil {
		t.Error("Expected validation error for invalid user, got nil")
	}
}
