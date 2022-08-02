package domain

import (
	"database/sql"
	"golang_lessons/custom_errors"
	"golang_lessons/logger"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TOKEN_DURATION = 10000

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       string         `db:"role"`
}

type LoginRepository interface {
	FindByUsername(string) (*Login, *custom_errors.AppErrors)
}

func (l Login) GenerateToken() (*string, *custom_errors.AppErrors) {
	var claims jwt.MapClaims

	if l.Accounts.Valid && l.CustomerId.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedTokensAsString, err := token.SignedString([]byte("test"))
	if err != nil {
		logger.Error("Failed while signing token")
		return nil, custom_errors.NewServerError("failed while signing token")
	}
	return &signedTokensAsString, nil
}

func (l Login) claimsForUser() jwt.MapClaims {
	accounts := strings.Split(l.Accounts.String, ",")
	return jwt.MapClaims{
		"customer_id": l.CustomerId.String,
		"role":        l.Role,
		"username":    l.Username,
		"accounts":    accounts,
		"exp":         time.Now().Add(TOKEN_DURATION).Unix(),
	}
}

func (l Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"username": l.Username,
		"role":     l.Role,
		"exp":      time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
