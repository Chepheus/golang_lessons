package domain

import (
	"golang_lessons/custom_errors"
	"time"
)

type Transaction struct {
	Id        int
	AccountId int
	Amount    float64
	Type      string
	Date      time.Time
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *custom_errors.AppErrors)
}
