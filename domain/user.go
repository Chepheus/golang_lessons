package domain

import (
	"golang_lessons/custom_errors"
	"time"
)

type User struct {
	Username   string
	Password   string
	Role       string
	CustomerId int
	CreatedOn  time.Time
}

type UserRepository interface {
	FindByUsername(string) (*User, *custom_errors.AppErrors)
}
