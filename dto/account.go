package dto

type NewAccountRequest struct {
	CustomerId int
	Type       string
	Amount     float64
}

type NewAccountResponse struct {
	Id         int
	CustomerId int
	Type       string
	Amount     float64
}
