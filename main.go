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
	r.HandleFunc(`/contests/{contest}/tasks/{task:[\d+]}`, judge).Methods("POST")
	r.HandleFunc("/submissions/{id}", connectClient).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
