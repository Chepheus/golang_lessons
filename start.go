package main

import (
	"golang_lessons/domain/repository"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"golang_lessons/handler"
	"golang_lessons/logger"
	"golang_lessons/service"
	"net/http"
	"time"
)

type CurrentTime struct {
	CurrentTime time.Time `json:"current_time"`
}

func Start() {
	router := mux.NewRouter()

	//customerHandler := handler.CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepository := repository.NewCustomerRepositoryDb(getDbClient())
	customerHandler := handler.CustomerHandler{
		Service: service.NewCustomerService(customerRepository),
	}

	accountRepository := repository.NewAccountRepositoryDb(getDbClient())
	accountHandler := handler.AccountHandler{
		Service: service.NewAccountService(accountRepository),
	}

	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}/account", accountHandler.NewAccount).Methods(http.MethodPost)

	err := http.ListenAndServe("localhost:8011", router)
	if err != nil {
		logger.Error("Server cant listen localhost:8011")
	}
}

func getDbClient() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:test@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
