package repository

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/logger"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func (r AccountRepositoryDb) Save(account domain.Account) (*domain.Account, *custom_errors.AppErrors) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES(?, ?, ?, ?, ?)"

	result, err := r.db.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.Type, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error while creating new account!")
		return nil, custom_errors.NewServerError(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting ID!")
		return nil, custom_errors.NewServerError(err.Error())
	}

	account.Id = int(id)
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{db: dbClient}
}
