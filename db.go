package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type detailRowType struct {
	Title  string `json:"title"`
	Result int    `json:"result"`
	Time   int64  `json:"time"`
	Memory int64  `json:"memory"`
}

type detailType struct {
	Contest      string          `json:"contest"`
	ContestID    string          `json:"contest_id"`
	Task         string          `json:"task"`
	TaskNumber   string          `json:"task_number"`
	WholeResult  int             `json:"whole_result"`
	MaxTime      int64           `json:"max_time"`
	MaxMemory    int64           `json:"max_memory"`
	CompileError string          `json:"compile_error"`
	Details      []detailRowType `json:"details"`
}

type taskInfoType struct {
	Title     string `json:"title"`
	Problem   string `json:"problem"`
	TimeLimit int    `json:"time_limit"`
}

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

func registerSubmission(id string, contest string, task int, submission submissionType, wholeResult resultType, maxMemory int64, maxTime int64, description string, detail []detailRowType) {
	bytes, err := json.Marshal(&detail)
	if err != nil {
		log.Println(err)
		return
	}
	rows, err := db.Query("INSERT INTO `submissions`(`id`, `contest`, `task`, `lang`, `code`, `whole_result`, `max_memory`, `max_time`, `compile_error`, `details`) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", id, contest, task, submission.Lang, submission.Code, wholeResult, maxMemory, maxTime, description, bytes)
	if err != nil {
		log.Println(err)
		return
	}
	rows.Close()
}

func getSubmissionDetail(id string) detailType {
	var detail detailType
	var bytes []byte
	rows, err := db.Query("SELECT `contests`.`title` AS `contest`, contest AS `contest_id`, `tasks`.`title` AS `task`, `submissions`.`task` AS `task_number`, `whole_result`, `max_time`, `max_memory`, `compile_error`, `details` FROM `submissions` JOIN `contests` ON contests.id = submissions.contest JOIN `tasks` USING(`contest`) WHERE `submissions`.`id` = ?;", id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	if rows.Next() {
		if err := rows.Scan(
			&detail.Contest,
			&detail.ContestID,
			&detail.Task,
			&detail.TaskNumber,
			&detail.WholeResult,
			&detail.MaxTime,
			&detail.MaxMemory,
			&detail.CompileError,
			&bytes); err != nil {
			log.Println(err)
			return detail
		}
		if err := json.Unmarshal(bytes, &detail.Details); err != nil {
			log.Println(err)
			return detail
		}
	}
	return detail
}

func getTaskInfo(contest string, task int) taskInfoType {
	var taskInfo taskInfoType
	rows, err := db.Query("SELECT `title`, `problem`, `time_limit` FROM tasks WHERE `contest` = ? AND `task` = ?;", contest, task)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	if rows.Next() {
		if err := rows.Scan(&taskInfo.Title, &taskInfo.Problem, &taskInfo.TimeLimit); err != nil {
			log.Println(err)
			return taskInfo
		}
	}
	return taskInfo
}

func initDB() {
	_db, err := sql.Open("mysql", "root:abcdefgh@/judge")
	db = _db
	if err != nil {
		log.Fatalln(err)
	}
}
