package main

import (
	"golang_lessons/domain/repository"
	"golang_lessons/handler"
	"golang_lessons/logger"
	"golang_lessons/service"
	"net/http"

	"github.com/gorilla/mux"
)

func StartAuth() {
	router := mux.NewRouter()

	db := getDbClient()
	userHandler := handler.UserHandler{Service: service.NewUserService(
		repository.NewUserRepository(db),
		repository.NewLoginRepository(db),
	)}

	router.HandleFunc("/auth/login", userHandler.Login).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8011", router)
	if err != nil {
		logger.Error("Server cant listen localhost:8011")
	}
}
