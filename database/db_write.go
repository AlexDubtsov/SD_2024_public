package database

import (
	"fmt"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func Basic_WriteToDB(slice_SinglePerson []structures.Basic_SinglePerson) string {
	var errStr string
	// Insert records to DB
	for i := range slice_SinglePerson {

		// Start a DB transaction
		tx, err := DB.Begin()
		if err != nil {
			return err.Error()
		}

		// Insert into DB
		_, err = tx.Exec("INSERT INTO BASIC (ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES, COMMENT) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", slice_SinglePerson[i].ID, slice_SinglePerson[i].Email, slice_SinglePerson[i].Name, slice_SinglePerson[i].Gender, slice_SinglePerson[i].Phone, slice_SinglePerson[i].BageID, slice_SinglePerson[i].EventID, slice_SinglePerson[i].EventName, slice_SinglePerson[i].EventDate, slice_SinglePerson[i].Created, slice_SinglePerson[i].Likes, slice_SinglePerson[i].Comment)
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

	return errStr
}
