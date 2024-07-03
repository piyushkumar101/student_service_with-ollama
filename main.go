package main

import (
	"log"
	"net/http"
	"student_service/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/students", controllers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", controllers.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", controllers.GetStudentByID).Methods("GET")
	r.HandleFunc("/students/{id}", controllers.UpdateStudentByID).Methods("PUT")
	r.HandleFunc("/students/{id}", controllers.DeleteStudentByID).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", controllers.GenerateStudentSummary).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
