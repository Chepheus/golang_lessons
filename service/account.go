package service

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/dto"
	"golang_lessons/validator"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *custom_errors.AppErrors)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *custom_errors.AppErrors) {
	validationError := validator.ValidateAccount(request)
	if validationError != nil {
		return nil, validationError
	}

	a := domain.Account{
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now(), //Format("2006-01-02 15:04:05"),
		Type:        request.Type,
		Amount:      request.Amount,
		Status:      1,
	}

	account, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	resp := dto.NewAccountResponse{
		Id:         account.Id,
		CustomerId: account.CustomerId,
		Type:       account.Type,
		Amount:     account.Amount,
	}

	return &resp, nil
}

func NewAccountService(repo domain.AccountRepository) AccountService {
	return DefaultAccountService{
		repo: repo,
	}
}
