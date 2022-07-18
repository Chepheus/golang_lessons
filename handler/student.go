package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Student struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
	Sex  string `json:"sex"`
}

var students = []Student{
	{ID: 1, Name: "Anton", Age: 27, City: "Kiyv", Sex: "M"},
	{ID: 2, Name: "Julia", Age: 23, City: "Mykolaiv", Sex: "F"},
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(students)
	if err != nil {
		fmt.Println("Server can't write to Writer")
	}
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId, _ := strconv.Atoi(vars["id"])

	var student Student
	for _, s := range students {
		if s.ID == studentId {
			student = s
		}
	}

	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(student)
	if err != nil {
		fmt.Println("Server can't write to Writer")
	}
}

func SetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST")
}
