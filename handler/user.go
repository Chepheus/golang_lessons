package handler

import (
	"encoding/json"
	"golang_lessons/dto"
	"golang_lessons/service"
	"net/http"
)

type UserHandler struct {
	Service service.UserService
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		encodeResponse(w, http.StatusInternalServerError, err)
		return
	}

	token, custErr := h.Service.Login(req)
	if custErr != nil {
		encodeResponse(w, http.StatusUnauthorized, err)
		return
	}
	encodeResponse(w, http.StatusOK, token)
}
