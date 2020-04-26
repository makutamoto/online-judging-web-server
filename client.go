package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func connectClient(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	vars := mux.Vars(r)
	if _, ok := judgingSubmissions[vars["id"]]; ok {
		judgingSubmissions[vars["id"]] = append(judgingSubmissions[vars["id"]], conn)
	} else {
		defer conn.Close()
		status := getSubmissionOverview(vars["id"])
		if err := conn.WriteJSON(status); err != nil {
			log.Println(err)
		}
	}
}
