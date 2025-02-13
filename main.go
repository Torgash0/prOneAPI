package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	task   string
	taskMu sync.Mutex
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	taskMu.Lock()
	defer taskMu.Unlock()
	fmt.Fprintf(w, "Hello, %s!", task)
}

func SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Task string `json:"task"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskMu.Lock()
	defer taskMu.Unlock()
	task = req.Task
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/set-task", SetTaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
