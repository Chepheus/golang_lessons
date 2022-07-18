package domain

import "golang_lessons/custom_errors"

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (r CustomerRepositoryStub) FindAll(filter *Filter) ([]Customer, *custom_errors.AppErrors) {
	return r.Customers, nil
}

func (r CustomerRepositoryStub) FindById(id int) (*Customer, *custom_errors.AppErrors) {
	for _, c := range r.Customers {
		if id == c.Id {
			return &c, nil
		}
	}

	return nil, nil
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{Id: 1, Name: "Anton", City: "Warshaw", Status: 1},
		{Id: 1, Name: "Dmytro", City: "Kyiv", Status: 1},
		{Id: 1, Name: "Artem", City: "Kyiv", Status: 1},
	}

	return CustomerRepositoryStub{Customers: customers}
}
