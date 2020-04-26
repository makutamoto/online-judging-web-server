package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type submissionType struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}

var judgingSubmissions map[string][]*websocket.Conn

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // delete.
}

func judge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var submission submissionType
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		if err := json.Unmarshal(bytes, &submission); err != nil {
			log.Println(err)
			return
		}
		uuid := uuid.New().String()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{ \"id\": \"%s\" }", uuid)
		vars := mux.Vars(r)
		task, _ := strconv.Atoi(vars["task"])
		bytes = prepareJSON(vars["contest"], task, submission)
		go sendData(bytes, uuid)
	}
}
