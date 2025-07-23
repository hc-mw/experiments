package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var (
	nextId = 1
	db     = make(map[int]Task)
	dbMu   sync.Mutex
)

func main() {
	http.HandleFunc("/tasks", TasksHandler)
	http.HandleFunc("/tasks/", TaskHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTasks(w, r)
	case http.MethodPost:
		PostTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/tasks/"):])

	switch r.Method {
	case http.MethodGet:
		GetTasks(w, r, id)
	case http.MethodPut:
		PutTask(w, r)
	case http.MethodDelete:
		DeleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

}

func PostTask(w http.ResponseWriter, r *http.Request) {

}

func PutTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
