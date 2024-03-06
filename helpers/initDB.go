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
	processTable, err := DB.Prepare(`
    CREATE TABLE if not exists PROCESSING(
		ID INTEGER PRIMARY KEY,
        EMAIL TEXT,
		NAME TEXT,
		GENDER TEXT,
		PHONE TEXT,
		BAGE_ID_AT_EVENT TEXT,
		EVENT_ID TEXT,
		EVENT_NAME TEXT,
		EVENT_DATE TEXT,
		DATE_CREATED TEXT,
		LIKES TEXT
    )
	`)
	if err != nil {
		log.Fatal(err)
	}
	processTable.Exec()

	// Create sql table if it does not exist
	participTable, err := DB.Prepare(`
    CREATE TABLE if not exists PARTICIPANTS(
        EMAIL TEXT PRIMARY KEY,
		NAME TEXT,
		GENDER TEXT,
		PHONE TEXT,
		CONTACT TEXT,
		AGE TEXT,
		AGE_GROUP TEXT,
		SUBMISSION_FORM_TIME TEXT,
		COMMENT TEXT,
		CREATED TEXT
    )
	`)
	if err != nil {
		log.Fatal(err)
	}
	participTable.Exec()

	// Create sql table if it does not exist
	eventsTable, err := DB.Prepare(`
    CREATE TABLE if not exists EVENTS(
        ID TEXT PRIMARY KEY,
		NAME TEXT,
		TIME TEXT,
		SPOT TEXT,
		AGE_GROUP TEXT,
		COMMENT TEXT,
		CREATED TEXT
    )
	`)
	if err != nil {
		log.Fatal(err)
	}
	eventsTable.Exec()

	// Defer the closing of the database connection
	defer DB.Close()
}
