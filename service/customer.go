package service

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
)

type CustomerService interface {
	GetAllCustomers(filter *domain.Filter) ([]domain.Customer, *custom_errors.AppErrors)
	GetCustomer(int) (*domain.Customer, *custom_errors.AppErrors)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(filter *domain.Filter) ([]domain.Customer, *custom_errors.AppErrors) {
	return s.repo.FindAll(filter)
}

func (s DefaultCustomerService) GetCustomer(id int) (*domain.Customer, *custom_errors.AppErrors) {
	return s.repo.FindById(id)
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repo}
}
