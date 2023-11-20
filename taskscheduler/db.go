package taskscheduler

import (
	"github.com/jmoiron/sqlx"
)

// Database stores all tasks added to our scheduler, for reference and recovery
type Database struct {
	conn *sqlx.Conn
}

func createDB() *Database {
	sqlx.Open("https://ordered-set-simenghe.turso.io/?authToken=[your-auth-token]")
	return &Database{nil}
}
