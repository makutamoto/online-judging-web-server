package main

import (
	"encoding/json"
	"io/ioutil"
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

func updateContestOverview(w http.ResponseWriter, r *http.Request) {
	var info contestInfoType
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, &info)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	err = updateContestOverviewDB(vars["contest"], info)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func updateContestExplanation(w http.ResponseWriter, r *http.Request) {
	var info contestInfoType
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, &info)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	err = updateContestExplanationDB(vars["contest"], info)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
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

func updateSystemOverview(w http.ResponseWriter, r *http.Request) {
	var overview systemOverviewType
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, &overview)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = updateSystemOverviewDB(overview)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
