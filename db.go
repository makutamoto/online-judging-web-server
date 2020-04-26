package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getTestData(contest string, task int) problemType {
	var problem problemType
	rows, err := db.Query("SELECT `time_limit`, `accuracy`, `testcases` FROM tasks WHERE `contest` = ? AND `task` = ?;", contest, task)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	if rows.Next() {
		var bytes []byte
		if err := rows.Scan(&problem.Limit, &problem.Accuracy, &bytes); err != nil {
			log.Println(err)
			return problem
		}
		if err := json.Unmarshal(bytes, &problem.Tests); err != nil {
			log.Println(err)
			return problem
		}
	}
	return problem
}

func registerSubmission(id string, contest string, task int, submission submissionType, wholeResult resultType, maxMemory int64, maxTime int64) {
	rows, err := db.Query("INSERT INTO `submissions`(id,contest, task, lang, code, whole_result, max_memory, max_time) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", id, contest, task, submission.Lang, submission.Code, wholeResult, maxMemory, maxTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows.Close()
}

func getSubmissionOverview(id string) statusType {
	var status statusType
	rows, err := db.Query("SELECT `whole_result`, `max_memory`, `max_time` FROM `submissions` WHERE `id` = ?;", id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	if rows.Next() {
		if err := rows.Scan(&status.Result, &status.Memory, &status.Time); err != nil {
			log.Println(err)
			return status
		}
		status.WholeResult = status.Result
	}
	return status
}

func initDB() {
	_db, err := sql.Open("mysql", "root:abcdefgh@/judge")
	db = _db
	if err != nil {
		log.Fatalln(err)
	}
}
