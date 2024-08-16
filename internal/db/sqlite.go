package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./task.db")
	if err != nil {
		return nil, err
	}
	taskStmt := "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, is_done BOOLEAN)"
	if _, err = db.Exec(taskStmt); err != nil {
		return nil, err
	}
	return db, nil
}
