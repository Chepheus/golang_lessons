package repository

import (
	"database/sql"
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/logger"

	"github.com/jmoiron/sqlx"
)

type LoginRepositoryDb struct {
	db *sqlx.DB
}

func (r LoginRepositoryDb) FindByUsername(username string) (*domain.Login, *custom_errors.AppErrors) {
	var l domain.Login

	sqlString := "SELECT username, customer_id, accounts, role FROM users WHERE customer_username = ?"
	err := r.db.Get(&l, sqlString, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.NewNotFoundError("Customer not found")
		}

		logger.Error(err.Error())
		return nil, custom_errors.NewServerError("Database error")
	}

	return &l, nil
}

func NewLoginRepository(db *sqlx.DB) LoginRepositoryDb {
	return LoginRepositoryDb{
		db: db,
	}
}
