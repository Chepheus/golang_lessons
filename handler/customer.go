package handler

import (
	"golang_lessons/custom_errors"
	"golang_lessons/dto"
	"golang_lessons/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (h CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	filter, customError := h._getFilter(r)
	if customError != nil {
		encodeResponse(w, customError.Code, customError.AsMessage())
		return
	}

	customers, customError := h.Service.GetAllCustomers(filter)
	if customError != nil {
		encodeResponse(w, customError.Code, customError.AsMessage())
		return
	}

	encodeResponse(w, http.StatusOK, customers)
}

func (h CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	customer, customError := h.Service.GetCustomer(id)
	if customError != nil {
		encodeResponse(w, customError.Code, customError.AsMessage())
		return
	}

	encodeResponse(w, http.StatusOK, customer)
}

func (h CustomerHandler) _getFilter(r *http.Request) (*dto.Filter, *custom_errors.AppErrors) {
	status := r.URL.Query().Get("status")

	if status != "" {
		return dto.NewFilterDTO(status)
	}

	return nil, nil
}
