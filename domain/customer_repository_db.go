package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang_lessons/custom_errors"
	"golang_lessons/logger"
	"time"
)

type CustomerRepositoryDB struct {
	db *sql.DB
}

func (r CustomerRepositoryDB) FindAll(filter *Filter) ([]Customer, *custom_errors.AppErrors) {
	sqlString := "SELECT customer_id, name, city, status FROM customers"

	rows, err := r._getRows(sqlString, filter)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.NewNotFoundError("Customers not found")
		}

		logger.Error(err.Error())
		return nil, custom_errors.NewServerError("Database error")
	}

	return r._hydrate(rows)
}

func (r CustomerRepositoryDB) FindById(id int) (*Customer, *custom_errors.AppErrors) {
	row := r.db.QueryRow("SELECT customer_id, name, city, status FROM customers WHERE customer_id = ?", id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.NewNotFoundError("Customer not found")
		}

		logger.Error(err.Error())
		return nil, custom_errors.NewServerError("Database error")
	}

	return &c, nil
}

func (r CustomerRepositoryDB) _getRows(sqlString string, filter *Filter) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	if filter != nil {
		sqlString += " WHERE status = ?"
		rows, err = r.db.Query(sqlString, filter.status)
	} else {
		rows, err = r.db.Query(sqlString)
	}

	return rows, err
}

func (r CustomerRepositoryDB) _hydrate(rows *sql.Rows) ([]Customer, *custom_errors.AppErrors) {
	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Status)
		if err != nil {
			logger.Error(err.Error())
			return nil, custom_errors.NewServerError(err.Error())
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	db, err := sql.Open("mysql", "root:test@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDB{db: db}
}
