package domain

import (
	"golang_lessons/custom_errors"
	"time"
)

type Account struct {
	Id          int
	CustomerId  int
	OpeningDate time.Time
	Type        string
	Amount      float64
	Status      int
}

type AccountRepository interface {
	Save(Account) (*Account, *custom_errors.AppErrors)
}
