package validator

import (
	"golang_lessons/custom_errors"
	"golang_lessons/dto"
)

var allowedTypes = map[string]int{
	"deposit":    1,
	"withdrawal": 1,
}

func ValidateAccount(request dto.NewAccountRequest) *custom_errors.AppErrors {
	_, ok := allowedTypes[request.Type]
	if !ok {
		return custom_errors.NewValidationError("Incorrect type of account")
	}

	if request.Amount <= 0 {
		return custom_errors.NewValidationError("Amount should be 1 or more")
	}

	return nil
}

func ValidateTransaction(request dto.TransactionRequest) *custom_errors.AppErrors {
	_, ok := allowedTypes[request.Type]
	if !ok {
		return custom_errors.NewValidationError("Incorrect type of transaction")
	}

	if request.Amount <= 0 {
		return custom_errors.NewValidationError("Amount should be 1 or more")
	}

	return nil
}
