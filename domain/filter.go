package domain

import "golang_lessons/custom_errors"

type Filter struct {
	status int
}

var statusMap = map[string]int{
	"active":   1,
	"inactive": 0,
}

func NewFilter(status string) (*Filter, *custom_errors.AppErrors) {
	s, ok := statusMap[status]

	if !ok {
		return nil, custom_errors.NewServerError("Status isn't support")
	}

	return &Filter{status: s}, nil
}
