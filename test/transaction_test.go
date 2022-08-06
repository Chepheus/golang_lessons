package test

import (
	"golang_lessons/dto"
	"golang_lessons/validator"
	"net/http"
	"testing"
)

func Test_should_return_error_when_type_is_invalid(t *testing.T) {
	tr := dto.TransactionRequest{Type: "invalid type"}

	err := validator.ValidateTransaction(tr)
	if err == nil {
		t.Error("Error not returned from ValidateTransaction")
	}

	if err.Message != "Incorrect type of transaction" {
		t.Error("Invalid message from ValidateTransaction")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code from ValidateTransaction")
	}
}

func Test_should_return_error_when_amount_less_then_zero(t *testing.T) {
	tr := dto.TransactionRequest{Type: "deposit", Amount: -1}

	err := validator.ValidateTransaction(tr)
	if err == nil {
		t.Error("Error not returned from ValidateTransaction")
	}

	if err.Message != "Amount should be 1 or more" {
		t.Error("Invalid message from ValidateTransaction")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code from ValidateTransaction")
	}
}
