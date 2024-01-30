package main

import (
	"database/sql"
	"fmt"

	"github.com/andersonprante/go_packaging/internal/child"
	"github.com/andersonprante/go_packaging/internal/logging"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Initialize logger with database connection
	logging.DefaultLogger = &logging.Logger{
		Prefix:   "INFO",
		Database: db,
	}

	// Use logger
	logging.DefaultLogger.Log("Hello, logging!")
	child.SomeFunction()
	fmt.Println("oi mundo")
}
