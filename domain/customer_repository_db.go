package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang_lessons/custom_errors"
	"golang_lessons/logger"
	"time"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func (r CustomerRepositoryDB) FindAll(filter *Filter) ([]Customer, *custom_errors.AppErrors) {
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

func (r CustomerRepositoryDB) FindById(id int) (*Customer, *custom_errors.AppErrors) {
	var c Customer

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

func (r CustomerRepositoryDB) _getRows(sqlString string, filter *Filter) ([]Customer, error) {
	var err error
	var customers []Customer

	if filter != nil {
		sqlString += " WHERE status = ?"
		err = r.db.Select(&customers, sqlString, filter.Status)
	} else {
		err = r.db.Select(&customers, sqlString)
	}

	return customers, err
}

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	db, err := sqlx.Open("mysql", "root:test@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDB{db: db}
}
