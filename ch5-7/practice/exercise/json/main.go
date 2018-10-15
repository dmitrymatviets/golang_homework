package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Worker struct {
	Id         int64
	Name       string
	Occupation string
	Salary     int64
}

const filePath = "./ch5-7/practice/exercise/json/json.json"
const dbPath = "./ch5-7/practice/exercise/json/db.db"

func main() {
	db, err := connect()
	defer db.Close()
	workers := getWorkersFromJson()

	for _, worker := range workers {
		_, err := db.Exec("INSERT OR REPLACE INTO `users` values(?,?,?,?)", worker.Id, worker.Name, worker.Occupation, worker.Salary)
		if err != nil {
			panic(err)
		}
	}

	db.Exec("INSERT OR REPLACE INTO `users` values(?,?,?,?)", 3, "Петр", "аутсорсер", 10000)

	rows, err := db.Query("SELECT * from users")
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("Из БД")
	for rows.Next() {
		var (
			id         int64
			name       string
			occupation string
			salary     int64
		)
		if err := rows.Scan(&id, &name, &occupation, &salary); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d. %s | %s | %d руб.\r\n", id, name, occupation, salary)
	}

}

func getWorkersFromJson() []Worker {

	filePath, _ := filepath.Abs(filePath)
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}
	var workers []Worker
	json.Unmarshal(bytes, &workers)
	return workers
}

func connect() (*sql.DB, error) {
	dbPath, _ := filepath.Abs(dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	return db, err
}
