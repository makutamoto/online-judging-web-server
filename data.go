package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type statusType struct {
	WholeResult int    `json:"whole_result"`
	Result      int    `json:"result"`
	Time        int64  `json:"time"`
	Memory      int64  `json:"memory"`
	CurrentCase int    `json:"current_case"`
	WholeCase   int    `json:"whole_case"`
	Description string `json:"description"`
}

type testType struct {
	In  string `json:"in"`
	Out string `json:"out"`
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

func sendData(data []byte, id string) {
	var result resultType
	var status statusType
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
		res := resultType(status.Result)
		fmt.Printf("%d/%d %v %dms %dkb\n", status.CurrentCase, status.WholeCase, res, status.Time, status.Memory)
		result.update(res)
		for _, connClient := range judgingSubmissions[id] {
			connClient.WriteMessage(websocket.TextMessage, bytes)
		}
	}
	fmt.Println(result)
	for _, connClient := range judgingSubmissions[id] {
		connClient.Close()
	}
	if resultType(status.WholeResult) == resultCompileError {
		fmt.Println(status.Description)
	}
}

func prepareJSON(submission submissionType) []byte {
	data := serverType{Language: submission.Lang, Code: submission.Code, Problem: problemType{
		Limit: 2000, Accuracy: -3, Tests: []testType{{In: "3 \n1 2 3\n", Out: "1\n2\nFizz\n"}, {In: "3\n2 3 1\n", Out: "2\nFizz\n1\n"}},
	}}
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return json
}
