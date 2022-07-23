package repository

import (
	"database/sql"
	"golang_lessons/custom_errors"
	"golang_lessons/domain"
	"golang_lessons/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func (r CustomerRepositoryDB) FindAll(filter *domain.Filter) ([]domain.Customer, *custom_errors.AppErrors) {
	sqlString := "SELECT customer_id, name, city, status FROM customers"

	customers, err := r._getRows(sqlString, filter)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.NewNotFoundError("Customers not found")
		}

		logger.Error(err.Error())
		return nil, custom_errors.NewServerError("Database error")
	}

	return customers, nil
}

func (r CustomerRepositoryDB) FindById(id int) (*domain.Customer, *custom_errors.AppErrors) {
	var c domain.Customer

	sqlString := "SELECT customer_id, name, city, status FROM customers WHERE customer_id = ?"
	err := r.db.Get(&c, sqlString, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.NewNotFoundError("Customer not found")
		}

		logger.Error(err.Error())
		return nil, custom_errors.NewServerError("Database error")
	}

	return &c, nil
}

func (r CustomerRepositoryDB) _getRows(sqlString string, filter *domain.Filter) ([]domain.Customer, error) {
	var err error
	var customers []domain.Customer

	if filter != nil {
		sqlString += " WHERE status = ?"
		err = r.db.Select(&customers, sqlString, filter.Status)
	} else {
		err = r.db.Select(&customers, sqlString)
	}

	return customers, err
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{db: dbClient}
}
