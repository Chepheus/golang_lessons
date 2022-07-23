package repository

import (
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/logger"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDb struct {
	db *sqlx.DB
}

func (r TransactionRepositoryDb) Save(transaction domain.Transaction) (*domain.Transaction, *custom_errors.AppErrors) {
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES(?, ?, ?, ?)"

	result, err := r.db.Exec(sqlInsert, transaction.AccountId, transaction.Amount, transaction.Type, transaction.Date)
	if err != nil {
		logger.Error("Error while creating new transaction!")
		return nil, custom_errors.NewServerError(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting ID!")
		return nil, custom_errors.NewServerError(err.Error())
	}

	transaction.Id = int(id)
	return &transaction, nil
}

func NewTransactionRepositoryDb(db *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{db: db}
}
