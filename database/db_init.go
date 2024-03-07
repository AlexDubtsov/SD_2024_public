package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Set the database file path
var dbAddress string = "../SD_2024_private/sd_database.db"
var DB *sql.DB

// Initialize the database connection
func init() {
	// Initialize the database connection
	DBopen()

	// Create the table if it doesn't exist
	if DB != nil {
		DBcreate()
	}
}

func DBopen() {
	var err error
	DB, err = sql.Open("sqlite3", dbAddress)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
}

func DBcreate() {
	// Create sql table if it does not exist
	basicTable, err := DB.Prepare(`
    CREATE TABLE if not exists BASIC(
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
		LIKES TEXT,
		COMMENT TEXT
    )
	`)
	if err != nil {
		log.Fatal(err)
	}
	basicTable.Exec()
}
