package main

import (
	"database/sql"
	"fmt"
	"log"

	mysql "github.com/go-sql-driver/mysql"
)

type Task struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
}

type TaskStore struct {
	DB *sql.DB
}

func connectDB() (TaskStore, error) {
	var store TaskStore

	cfg := mysql.Config{
		User:   "root",
		Passwd: "mysql@123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "goapi",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return store, err
	}

	if pingErr := db.Ping(); pingErr != nil {
		return store, err
	}

	fmt.Println("Connected !")
	store.DB = db

	return store, nil
}

func (ts *TaskStore) GetAllTasks() ([]Task, error) {
	var tasks []Task

	rows, err := ts.DB.Query("select * from tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	taskStore, err := connectDB()
	if err != nil {
		panic(err)
	}

	tasks, err := taskStore.GetAllTasks()
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		fmt.Println("ID: ", task.Id, ", Title: ", task.Title)
	}
}
