package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	initDB()
	defer db.Close()
	judgingSubmissions = map[string][]*websocket.Conn{}
	r := mux.NewRouter()
	r.HandleFunc(`/api/`, getSystemOverview).Methods("GET")
	r.HandleFunc(`/api/`, updateSystemOverview).Methods("PUT")
	r.HandleFunc(`/api/contests/`, getContestList).Methods("GET")
	r.HandleFunc(`/api/contests/{contest}`, getContestInfo).Methods("GET")
	r.HandleFunc(`/api/contests/{contest}`, updateContestOverview).Methods("PUT")
	r.HandleFunc(`/api/contests/{contest}/explanation`, updateContestExplanation).Methods("PUT")
	r.HandleFunc(`/api/contests/{contest}/tasks/`, getTaskList).Methods("GET")
	r.HandleFunc(`/api/contests/{contest}/tasks/{task:[\d+]}`, getTaskInfo).Methods("GET")
	r.HandleFunc(`/api/contests/{contest}/tasks/{task:[\d+]}`, updateTaskProblem).Methods("PUT")
	r.HandleFunc(`/api/contests/{contest}/tasks/{task:[\d+]}`, judge).Methods("POST")
	r.HandleFunc("/api/submissions/realtime/{id}", getRealtime).Methods("GET")
	r.HandleFunc("/api/submissions/details/{id}", getSubmissionDetail).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
