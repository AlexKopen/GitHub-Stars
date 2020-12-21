package main

import "testing"

func TestInputValidation(t *testing.T) {
	input := "angular/angular"
	validInputs := generateValidInputs(input, false)

	if len(validInputs) != 1 {
		t.Fatalf("Expected the input of '%s' to return 1 valid input, but instead received %d valid inputs.", input, len(validInputs))
	}
}

func TestMultipleInputValidation(t *testing.T) {
	input := "angular/angular, twilio/twilio-python"
	validInputs := generateValidInputs(input, false)

	if len(validInputs) != 2 {
		t.Fatalf("Expected the input of '%s' to return 2 valid inputs, but instead received %d valid inputs.", input, len(validInputs))
	}
}

func TestNoSlashInputValidation(t *testing.T) {
	input := "angular"
	validInputs := generateValidInputs(input, false)

	if len(validInputs) != 0 {
		t.Fatalf("Expected the input of '%s' to return 0 valid inputs, but instead received %d valid inputs.", input, len(validInputs))
	}
}

func TestInvalidRepoInputValidation(t *testing.T) {
	input := "angular/"
	validInputs := generateValidInputs(input, false)

	if len(validInputs) != 0 {
		t.Fatalf("Expected the input of '%s' to return 0 valid inputs, but instead received %d valid inputs.", input, len(validInputs))
	}
}

func TestResultsFromServer(t *testing.T) {
	input := "angular/angular"
	validInputs := []string{"angular/angular"}
	output := processInput(validInputs)

	if len(output) == 0 {
		t.Fatalf("Expected the input of '%s' to generate a result from the server, but instead no results were returned.", input)
	}
}
