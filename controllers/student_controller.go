package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"student_service/models"
	"student_service/ollama"
	"student_service/views"

	"github.com/gorilla/mux"
)

var (
	students  = make(map[int]models.Student)
	idCounter = 1
	mu        sync.Mutex
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	student.ID = idCounter
	idCounter++
	students[student.ID] = student
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	views.JSONResponse(w, student)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var studentList []models.Student
	for _, student := range students {
		studentList = append(studentList, student)
	}

	views.JSONResponse(w, studentList)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	student, exists := students[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	views.JSONResponse(w, student)
}

func UpdateStudentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedStudent models.Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	student, exists := students[id]
	if !exists {
		mu.Unlock()
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	student.Name = updatedStudent.Name
	student.Age = updatedStudent.Age
	student.Email = updatedStudent.Email
	students[id] = student
	mu.Unlock()

	views.JSONResponse(w, student)
}

func DeleteStudentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	_, exists := students[id]
	if !exists {
		mu.Unlock()
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	delete(students, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func GenerateStudentSummary(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	student, exists := students[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	summary, err := ollama.FetchStudentSummaryFromOllama(student)

	fmt.Printf("Error is %v", err)
	if err != nil {
		http.Error(w, "Error generating summary", http.StatusInternalServerError)
		return
	}

	views.JSONResponse(w, map[string]string{"summary": summary})
}
