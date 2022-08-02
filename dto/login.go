package dto

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username  string
	CreatedOn time.Time
}
