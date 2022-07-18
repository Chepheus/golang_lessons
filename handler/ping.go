package handler

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "PONG!")
	if err != nil {
		fmt.Println("Server can't write to Writer")
	}
}
