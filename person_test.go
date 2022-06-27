package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPerson(t *testing.T) {
	p, err := NewPerson("Kai", "Clairvoyant", 28)
	assert.Equal(t, p.name, "Kai")
	assert.Equal(t, p.lastName, "Clairvoyant")
	assert.Equal(t, p.age, 28)
	assert.Nil(t, err)
}

func TestValidations(t *testing.T) {
	standardInput := Person{"Kai", "Clairvoyant", 28}

	// map["value"] = "errorMessage"
	nameErrors := make(map[string]string)
	lastNameErrors := make(map[string]string)
	ageErrors := make(map[int]string)

	// name validations
	nameErrors["N"] = "name is too short"
	nameErrors["invalid"] = "name is invalid, literally"

	// last name validations
	lastNameErrors["L"] = "last name is too short"

	// age validations
	ageErrors[14] = "user is under the age of majority"
	ageErrors[200] = "no vampires allowed"

	for value, message := range nameErrors {
		input := standardInput
		input.name = value
		assertValidationReturnsTheCorrectErrorMessage(t, input, message)
	}

	for value, message := range lastNameErrors {
		input := standardInput
		input.lastName = value
		assertValidationReturnsTheCorrectErrorMessage(t, input, message)
	}

	for value, message := range ageErrors {
		input := standardInput
		input.age = value
		assertValidationReturnsTheCorrectErrorMessage(t, input, message)
	}
}

func assertValidationReturnsTheCorrectErrorMessage(t *testing.T, input Person, expectedMsg string) {
	p, err := NewPerson(input.name, input.lastName, input.age)
	assert.Nil(t, p)
	assert.Equal(t, errors.New(expectedMsg), err)
}
