package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getTestData(contest string, task int) problemType {
	var problem problemType
	rows, err := db.Query("SELECT `time_limit`, `accuracy`, `testcases` FROM tasks WHERE `contest_id` = ? AND `problem_number` = ?;", contest, task)
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
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

func initDB() {
	_db, err := sql.Open("mysql", "root:abcdefgh@/judge")
	db = _db
	if err != nil {
		log.Fatalln(err)
	}
}
