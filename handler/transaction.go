package handler

import (
	"encoding/json"
	"golang_lessons/dto"
	"golang_lessons/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	Service service.TransactionService
}

func (h TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["id"]

	var req dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		encodeResponse(w, http.StatusInternalServerError, err)
		return
	}

	req.AccountId, _ = strconv.Atoi(accountId)
	transaction, customErr := h.Service.NewTransaction(req)
	if customErr != nil {
		encodeResponse(w, customErr.Code, customErr.AsMessage())
		return
	}

	encodeResponse(w, http.StatusOK, transaction)
}
