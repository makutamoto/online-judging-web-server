package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type connectInfoType struct {
	ID string `json:"id"`
}

func connectClient(w http.ResponseWriter, r *http.Request) {
	var info connectInfoType
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	_, bytes, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(bytes, &info)
	if err != nil {
		log.Println(err)
		return
	}
	judgingSubmissions[info.ID] = append(judgingSubmissions[info.ID], conn)
}
