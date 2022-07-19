package service

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/dto"
)

type CustomerService interface {
	GetAllCustomers(filter *dto.Filter) ([]domain.Customer, *custom_errors.AppErrors)
	GetCustomer(int) (*dto.Customer, *custom_errors.AppErrors)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(filter *dto.Filter) ([]domain.Customer, *custom_errors.AppErrors) {
	if filter != nil {
		s.repo.FindAll(&domain.Filter{Status: filter.Status})
	}
	return s.repo.FindAll(nil)
}

func (s DefaultCustomerService) GetCustomer(id int) (*dto.Customer, *custom_errors.AppErrors) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewCustomerDTOFromDomain(c), nil
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repo}
}
