package helpers

import (
	"fmt"
	"os"
)

// Readout AllData, MaximumEventID and Maximum Record ID from Basic Processing DB
func BasicReadoutDB() ([]BasicSinglePerson, int, int) {
	var basicAllData []BasicSinglePerson
	var maxEventID int
	var maxID int

	// Initialize the database connection
	DBopen()

	// Query the database to get all records
	records, err := DB.Query("SELECT ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES FROM PROCESSING")
	if err != nil {
		fmt.Println("ERROR: Unable to read data base", err)
		os.Exit(1)
	}

	defer records.Close()
	// Defer the closing of the database connection
	defer DB.Close()

	// Iterate over the records; Add each record to RESULT
	for records.Next() {
		var record BasicSinglePerson
		err := records.Scan(&record.ID, &record.Email, &record.Name, &record.Gender, &record.Phone, &record.BageID, &record.EventID, &record.EventName, &record.EventDate, &record.Created, &record.Likes)
		if err != nil {
			fmt.Println("ERROR: Unable to scan records", err)
			os.Exit(1)
		}
		basicAllData = append(basicAllData, record)
		if record.EventID > maxEventID {
			maxEventID = record.EventID
		}
		if record.ID > maxID {
			maxID = record.ID
		}
	}
	return basicAllData, maxEventID, maxID
}

// Write NewData to Basic Processing DB
func BasicWriteDB(arrayBasicPersons []BasicSinglePerson) string {

	// Initialize the database connection
	DBopen()
	// Defer the closing of the database connection
	defer DB.Close()

	// Insert records to DB
	for _, element := range arrayBasicPersons {

		// Start a DB transaction
		tx, err := DB.Begin()
		if err != nil {
			// Rollback the transaction in case of an error
			DB.Close()
			return err.Error()
		}

		// Insert into DB
		_, err = tx.Exec("INSERT INTO PROCESSING (ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", element.ID, element.Email, element.Name, element.Gender, element.Phone, element.BageID, element.EventID, element.EventName, element.EventDate, element.Created, element.Likes)
		if err != nil {
			fmt.Println("Error inserting record:", err)
			// Rollback the transaction in case of an error
			tx.Rollback()
			return err.Error()
		}

		// Commit the DB transaction
		err = tx.Commit()
		if err != nil {
			fmt.Println("Error committing transaction:", err)
			return err.Error()
		}
	}

	return ""
}
