package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang_lessons/custom_errors"
	"golang_lessons/dto"
	"golang_lessons/service"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (h CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	filter, customError := h._getFilter(r)
	if customError != nil {
		h._encodeResponse(w, customError.Code, customError.AsMessage())
		return
	}

	customers, customError := h.Service.GetAllCustomers(filter)
	if customError != nil {
		h._encodeResponse(w, customError.Code, customError.AsMessage())
		return
	}

	h._encodeResponse(w, http.StatusOK, customers)
}

func (h CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	customer, customError := h.Service.GetCustomer(id)
	if customError != nil {
		h._encodeResponse(w, customError.Code, customError.AsMessage())
		return
	}

	h._encodeResponse(w, http.StatusOK, customer)
}

func (h CustomerHandler) _getFilter(r *http.Request) (*dto.Filter, *custom_errors.AppErrors) {
	status := r.URL.Query().Get("status")

	if status != "" {
		return dto.NewFilterDTO(status)
	}

	return nil, nil
}

func (h CustomerHandler) _encodeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
