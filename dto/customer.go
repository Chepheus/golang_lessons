package dto

import "golang_lessons/domain"

type Customer struct {
	Id     int    `json:"customer_id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Status int    `json:"status"`
}

func NewCustomerDTOFromDomain(c *domain.Customer) *Customer {
	return &Customer{
		Id:     c.Id,
		Name:   c.Name,
		City:   c.City,
		Status: c.Status,
	}
}
