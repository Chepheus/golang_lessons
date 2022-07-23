package service

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/dto"
	"golang_lessons/validator"
	"time"
)

type TransactionService interface {
	NewTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *custom_errors.AppErrors)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) NewTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *custom_errors.AppErrors) {
	validationError := validator.ValidateTransaction(request)
	if validationError != nil {
		return nil, validationError
	}

	t := domain.Transaction{
		AccountId: request.AccountId,
		Type:      request.Type,
		Amount:    request.Amount,
		Date:      time.Now(), //Format("2006-01-02 15:04:05"),
	}

	transaction, err := s.repo.Save(t)
	if err != nil {
		return nil, err
	}

	resp := dto.TransactionResponse{
		Id:        transaction.Id,
		AccountId: transaction.AccountId,
		Type:      transaction.Type,
		Amount:    transaction.Amount,
		Date:      transaction.Date,
	}

	return &resp, nil
}

func NewTransactionService(repo domain.TransactionRepository) TransactionService {
	return DefaultTransactionService{
		repo: repo,
	}
}
