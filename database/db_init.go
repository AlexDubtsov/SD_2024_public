package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/console"
	_ "github.com/mattn/go-sqlite3"
)

// Set the database file path
var dbAddress string = "../SD_2024_private/sd_database.db"
var DB *sql.DB

// Initialize the database connection
func init() {
	// Load latest DB from Cloud
	console.ConsoleLoad()

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
		fmt.Println("Error opening database")
		log.Fatal(err)
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
