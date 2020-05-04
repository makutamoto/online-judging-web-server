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

func getSubmissionDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	detail := getSubmissionDetailDB(vars["id"])
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

func getTaskInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	task, _ := strconv.Atoi(vars["task"])
	detail := getTaskInfoDB(vars["contest"], task)
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

func getTaskList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	list := getTaskListDB(vars["contest"])
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

func getContestInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	list := getContestInfoDB(vars["contest"])
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

func getContestList(w http.ResponseWriter, r *http.Request) {
	list := getContestListDB()
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

func getSystemOverview(w http.ResponseWriter, r *http.Request) {
	overview := getSystemOverviewDB()
	bytes, err := json.Marshal(&overview)
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
