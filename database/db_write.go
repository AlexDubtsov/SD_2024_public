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

func Basic_ChangeRecordsDB(eventID int, formName, formDate string) structures.Template_Basic_EditEvent {
	var templateData structures.Template_Basic_EditEvent
	// Start a DB transaction
	tx, err := DB.Begin()
	if err != nil {
		templateData.Message = err.Error()
		return templateData
	}

	// Query the database to get all transactions
	records, err := tx.Query("SELECT ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES, COMMENT FROM BASIC where EVENT_ID = ?", eventID)
	if err != nil {
		fmt.Println("Error in Query the database")
		templateData.Message = "Error in Query the database"
		return templateData
	}
	defer records.Close()

	// Iterate through the records, update the EVENT_NAME and EVENT_DATE
	for records.Next() {
		var record structures.Basic_SinglePerson

		err := records.Scan(&record.ID, &record.Email, &record.Name, &record.Gender, &record.Phone, &record.BageID, &record.EventID, &record.EventName, &record.EventDate, &record.Created, &record.Likes, &record.Comment)
		if err != nil {
			fmt.Println("Error scanning record")
			templateData.Message = "Error scanning record"
			return templateData
		}

		_, err = tx.Exec("UPDATE BASIC SET EVENT_NAME = ?, EVENT_DATE = ? WHERE EVENT_ID = ?", formName, formDate, eventID)
		if err != nil {
			fmt.Println("Error updating record")
			templateData.Message = "Error updating record"
			return templateData
		}
	}

	// Commit the DB transaction
	err = tx.Commit()
	if err != nil {
		templateData.Message = "Error on Commit to data base"
		return templateData
	}

	templateData.Message = "Updated Event Name, Event Date"
	return templateData
}
