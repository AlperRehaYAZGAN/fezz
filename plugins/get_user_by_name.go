//go:build plugin

package main

// go build -buildmode=plugin -ldflags="-s -w" -o plugins/get_user_by_name.so plugins/get_user_by_name.go

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbConn *sql.DB

// User represents a user record from the SQLite database.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handle handles JSON-RPC requests to fetch user records from the SQLite database.
func OnInit() {
	wd, err := os.Getwd()
	// Open a connection to the SQLite database
	dbPath := wd + "/db.sqlite3"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	dbConn = db

	// create temp table if not exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	// insert sample "john" user
	_, err = db.Exec("INSERT INTO users (name, age) VALUES ('john', 30)")
	if err != nil {
		panic(err)
	}
}

func OnKill() {
	dbConn.Close()
}

type Dto struct {
	Name string `json:"name"`
}

// Handle handles JSON-RPC requests to fetch user records from the SQLite database.
func Handle(params []interface{}) (string, error) {
	// marshall []interface first index to Dto struct
	name := "john"

	// Query the database for the first user record with the given name
	var user User
	err := dbConn.QueryRow("SELECT id, name, age FROM users WHERE name = ? LIMIT 1", name).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return "", err
	}

	// return &user, nil
	return "User: " + user.Name + ", Age: " + string(user.Age), nil
}
