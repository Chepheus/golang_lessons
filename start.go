package main

import (
	"github.com/gorilla/mux"
	"golang_lessons/domain"
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
	customerHandler := handler.CustomerHandler{
		Service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}
	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8011", router)
	if err != nil {
		logger.Error("Server cant listen localhost:8011")
	}
}
