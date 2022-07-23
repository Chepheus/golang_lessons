package handler

import (
	"encoding/json"
	"golang_lessons/dto"
	"golang_lessons/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	Service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	var req dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		encodeResponse(w, http.StatusInternalServerError, err)
		return
	}

	req.CustomerId, _ = strconv.Atoi(customerId)
	account, customErr := h.Service.NewAccount(req)
	if customErr != nil {
		encodeResponse(w, customErr.Code, customErr.AsMessage())
		return
	}

	encodeResponse(w, http.StatusOK, account)
}
