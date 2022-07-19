package dto

import (
	"golang_lessons/custom_errors"
)

type Filter struct {
	Status int
}

var statusMap = map[string]int{
	"active":   1,
	"inactive": 0,
}

func NewFilterDTO(status string) (*Filter, *custom_errors.AppErrors) {
	s, ok := statusMap[status]

	if !ok {
		return nil, custom_errors.NewServerError("Status isn't support")
	}

	return &Filter{Status: s}, nil
}
