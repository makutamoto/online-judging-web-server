package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getRealtime(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	vars := mux.Vars(r)
	if _, ok := judgingSubmissions[vars["id"]]; ok {
		judgingSubmissions[vars["id"]] = append(judgingSubmissions[vars["id"]], conn)
	} else {
		conn.Close()
	}
}

func getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	detail := getSubmissionDetail(vars["id"])
	bytes, err := json.Marshal(&detail)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
		return
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	task, _ := strconv.Atoi(vars["task"])
	detail := getTaskInfo(vars["contest"], task)
	bytes, err := json.Marshal(&detail)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
		return
	}
}

func getTaskOverviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	list := getTaskList(vars["contest"])
	bytes, err := json.Marshal(&list)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
		return
	}
}

func getContest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	list := getContestInfo(vars["contest"])
	bytes, err := json.Marshal(&list)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
		return
	}
}
