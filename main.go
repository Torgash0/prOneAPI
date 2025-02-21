package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := DB.Create(&newTask).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {

	InitDB()
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks", GetTasksHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
