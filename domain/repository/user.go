package repository

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	db *sqlx.DB
}

func (r UserRepositoryDb) FindByUsername(username string) (*domain.User, *custom_errors.AppErrors) {
	return nil, nil
}

func NewUserRepository(db *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{
		db: db,
	}
}
