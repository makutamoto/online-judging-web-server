package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type statusType struct {
	End         bool   `json:"end"`
	Title       string `json:"title"`
	WholeResult int    `json:"whole_result"`
	Result      int    `json:"result"`
	Time        int64  `json:"time"`
	Memory      int64  `json:"memory"`
	CurrentCase int    `json:"current_case"`
	WholeCase   int    `json:"whole_case"`
	Description string `json:"description"`
}

type testType struct {
	Title string `json:"title"`
	In    string `json:"in"`
	Out   string `json:"out"`
}

type problemType struct {
	Limit    int        `json:"limit"`
	Accuracy int        `json:"accuracy"`
	Tests    []testType `json:"tests"`
}

type serverType struct {
	Language string      `json:"language"`
	Code     string      `json:"code"`
	Problem  problemType `json:"problem"`
}

func sendData(id string, contest string, task int, submission submissionType) {
	var result resultType
	var status statusType
	var maxMemory, maxTime int64
	var detail []detailRowType
	data := prepareJSON(contest, task, submission)
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:7867/submit", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Fatalln(err)
	}
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if err := json.Unmarshal(bytes, &status); err != nil {
			log.Fatalln(err)
		}
		if maxMemory < status.Memory {
			maxMemory = status.Memory
		}
		if maxTime < status.Time {
			maxTime = status.Time
		}
		res := resultType(status.Result)
		fmt.Printf("%d/%d %v %dms %dkb\n", status.CurrentCase, status.WholeCase, res, status.Time, status.Memory)
		result.update(res)
		if status.End && res != resultCompileError {
			detail = append(detail, detailRowType{Title: status.Title, Result: status.Result, Time: status.Time, Memory: status.Memory})
		}
		for _, connClient := range judgingSubmissions[id] {
			connClient.WriteMessage(websocket.TextMessage, bytes)
		}
	}
	fmt.Println(result)
	registerSubmission(id, contest, task, submission, result, maxMemory, maxTime, status.Description, detail)
	for _, connClient := range judgingSubmissions[id] {
		connClient.Close()
	}
	delete(judgingSubmissions, id)
}

func prepareJSON(contest string, task int, submission submissionType) []byte {
	problem := getTestData("bc1", 1)
	data := serverType{Language: submission.Lang, Code: submission.Code, Problem: problem}
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return json
}
