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
	Lang         string          `json:"lang"`
	Code         string          `json:"code"`
	CompileError string          `json:"compile_error"`
	Details      []detailRowType `json:"details"`
}

type taskInfoType struct {
	Title     string `json:"title"`
	Problem   string `json:"problem"`
	TimeLimit int    `json:"time_limit"`
}

type taskOverviewType struct {
	Title     string `json:"title"`
	TimeLimit int    `json:"time_limit"`
}

type contestInfoType struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Explanation string `json:"explanation"`
}

type contestOverviewType struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type systemOverviewType struct {
	Overview string `json:"overview"`
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

func getSubmissionDetailDB(id string) detailType {
	var detail detailType
	var bytes []byte
	rows, err := db.Query("SELECT `contests`.`title` AS `contest`, contest AS `contest_id`, `tasks`.`title` AS `task`, `submissions`.`task` AS `task_number`, `whole_result`, `max_time`, `max_memory`, `lang`, `code`, `compile_error`, `details` FROM `submissions` JOIN `contests` ON contests.id = submissions.contest JOIN `tasks` USING(`contest`) WHERE `submissions`.`id` = ?;", id)
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
			&detail.Lang,
			&detail.Code,
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

func getTaskInfoDB(contest string, task int) taskInfoType {
	var taskInfo taskInfoType
	rows, err := db.Query("SELECT `title`, `problem`, `time_limit` FROM tasks WHERE `contest` = ? AND `task` = ?;", contest, task)
	if err != nil {
		log.Println(err)
		return taskInfo
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&taskInfo.Title, &taskInfo.Problem, &taskInfo.TimeLimit); err != nil {
			log.Println(err)
			return taskInfo
		}
	}
	return taskInfo
}

func getTaskListDB(contest string) []taskOverviewType {
	var taskList []taskOverviewType
	rows, err := db.Query("SELECT `title`, `time_limit` FROM `tasks` WHERE `contest` = ?;", contest)
	if err != nil {
		log.Println(err)
		return taskList
	}
	defer rows.Close()
	for rows.Next() {
		var overview taskOverviewType
		if err := rows.Scan(&overview.Title, &overview.TimeLimit); err != nil {
			log.Println(err)
			return taskList
		}
		taskList = append(taskList, overview)
	}
	return taskList
}

func getContestInfoDB(contest string) contestInfoType {
	var info contestInfoType
	rows, err := db.Query("SELECT `title`, `description`, `explanation` FROM `contests` WHERE `id` = ?;", contest)
	if err != nil {
		log.Println(err)
		return info
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&info.Title, &info.Description, &info.Explanation); err != nil {
			log.Println(err)
			return info
		}
	}
	return info
}

func getContestListDB() []contestOverviewType {
	var list []contestOverviewType
	rows, err := db.Query("SELECT `id`, `title` FROM `contests`;")
	if err != nil {
		log.Println(err)
		return list
	}
	defer rows.Close()
	for rows.Next() {
		var overview contestOverviewType
		if err := rows.Scan(&overview.ID, &overview.Title); err != nil {
			log.Println(err)
			return list
		}
		list = append(list, overview)
	}
	return list
}

func getSystemOverviewDB() systemOverviewType {
	var overview systemOverviewType
	rows, err := db.Query("SELECT `overview` FROM `system`;")
	if err != nil {
		log.Println(err)
		return overview
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&overview.Overview); err != nil {
			log.Println(err)
			return overview
		}
	}
	return overview
}

func updateSystemOverviewDB(overview systemOverviewType) error {
	_, err := db.Query("UPDATE `system` SET `overview` = ?;", overview.Overview)
	return err
}

func initDB() {
	_db, err := sql.Open("mysql", "root:abcdefgh@/judge")
	db = _db
	if err != nil {
		log.Fatalln(err)
	}
}
