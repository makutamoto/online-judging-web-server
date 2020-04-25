package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	judgingSubmissions = map[string][]*websocket.Conn{}
	http.HandleFunc("/submit", judge)
	http.HandleFunc("/submissions", connectClient)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
