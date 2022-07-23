package dto

import "time"

type TransactionRequest struct {
	AccountId int
	Amount    float64
	Type      string
	Date      time.Time
}

type TransactionResponse struct {
	Id        int
	AccountId int
	Amount    float64
	Type      string
	Date      time.Time
}
