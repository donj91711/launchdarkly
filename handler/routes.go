package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root path")
	})

	router.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		ResultCount(w)
	})

	router.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		Students(w)
	})

	router.HandleFunc("/students/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		studentID := vars["id"]
		Student(w, studentID)
	})

	router.HandleFunc("/exams", func(w http.ResponseWriter, r *http.Request) {
		Exams(w)
	})

	router.HandleFunc("/exams/{number}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		exam := vars["number"]
		Exam(w, exam)
	})
}
