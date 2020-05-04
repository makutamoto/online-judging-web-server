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
	r.HandleFunc(`/json`, getSystemOverview).Methods("GET")
	r.HandleFunc(`/contests/json`, getContestList).Methods("GET")
	r.HandleFunc(`/contests/{contest}/json`, getContestInfo).Methods("GET")
	r.HandleFunc(`/contests/{contest}/tasks/json`, getTaskList).Methods("GET")
	r.HandleFunc(`/contests/{contest}/tasks/{task:[\d+]}/json`, getTaskInfo).Methods("GET")
	r.HandleFunc(`/contests/{contest}/tasks/{task:[\d+]}`, judge).Methods("POST")
	r.HandleFunc("/submissions/realtime/{id}", getRealtime).Methods("GET")
	r.HandleFunc("/submissions/details/{id}", getSubmissionDetail).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
