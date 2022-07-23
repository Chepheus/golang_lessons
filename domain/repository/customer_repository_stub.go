package repository

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
)

type CustomerRepositoryStub struct {
	Customers []domain.Customer
}

func (r CustomerRepositoryStub) FindAll(filter *domain.Filter) ([]domain.Customer, *custom_errors.AppErrors) {
	return r.Customers, nil
}

func (r CustomerRepositoryStub) FindById(id int) (*domain.Customer, *custom_errors.AppErrors) {
	for _, c := range r.Customers {
		if id == c.Id {
			return &c, nil
		}
	}

	return nil, nil
}

func NewCustomerRepositoryStub() domain.CustomerRepository {
	customers := []domain.Customer{
		{Id: 1, Name: "Anton", City: "Warshaw", Status: 1},
		{Id: 1, Name: "Dmytro", City: "Kyiv", Status: 1},
		{Id: 1, Name: "Artem", City: "Kyiv", Status: 1},
	}

	return CustomerRepositoryStub{Customers: customers}
}
