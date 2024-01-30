package logging

import (
	"database/sql"
	"fmt"
)

type Logger struct {
	Prefix   string
	Database *sql.DB
}

func (l *Logger) Log(message string) {
	// Example: Insert log message into a database
	_, err := l.Database.Exec("INSERT INTO logs (prefix, message) VALUES (?, ?)", l.Prefix, message)
	if err != nil {
		fmt.Printf("Error logging message: %s\n", err)
		return
	}

	fmt.Printf("[%s] %s\n", l.Prefix, message)
}

// Shared instance
var DefaultLogger *Logger
