package domain

import "golang_lessons/custom_errors"

type Customer struct {
	Id     int
	Name   string
	City   string
	Status int
}

type CustomerRepository interface {
	FindAll(filter *Filter) ([]Customer, *custom_errors.AppErrors)
	FindById(int) (*Customer, *custom_errors.AppErrors)
}
