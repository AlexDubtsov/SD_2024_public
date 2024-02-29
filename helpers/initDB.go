package helpers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize the database connection
func DBopen() {
	var err error
	DB, err = sql.Open("sqlite3", "../SD_2024_private/sd_database.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
}

func DBcreate() {

	// Initialize the database connection
	DBopen()

	// Create sql table if it does not exist
	usersTable, err := DB.Prepare(`
    CREATE TABLE if not exists PARTICIPANTS(
        EMAIL TEXT PRIMARY KEY,
		NAME TEXT,
		PHONE TEXT,
		CREATED TEXT
    )
	`)
	if err != nil {
		log.Fatal(err)
	}
	usersTable.Exec()

	// Defer the closing of the database connection
	defer DB.Close()
}
