package database

import (
	"fmt"
	"log"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func Basic_WriteToDB(slice_SinglePerson []structures.Basic_SinglePerson) string {
	var errStr string
	// Insert records to DB
	for i := range slice_SinglePerson {

		// Start a DB transaction
		tx, err := DB.Begin()
		if err != nil {
			log.Fatal(err)
			return "Error on data base Begin"
		}

		// Insert into DB
		_, err = tx.Exec("INSERT INTO BASIC (ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES, COMMENT) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", slice_SinglePerson[i].ID, slice_SinglePerson[i].Email, slice_SinglePerson[i].Name, slice_SinglePerson[i].Gender, slice_SinglePerson[i].Phone, slice_SinglePerson[i].BageID, slice_SinglePerson[i].EventID, slice_SinglePerson[i].EventName, slice_SinglePerson[i].EventDate, slice_SinglePerson[i].Created, slice_SinglePerson[i].Likes, slice_SinglePerson[i].Comment)
		if err != nil {
			log.Fatal(err)
			// Rollback the transaction in case of an error
			tx.Rollback()
			return "Error inserting record"
		}

		// Commit the DB transaction
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
			fmt.Println("Error committing transaction")
			return err.Error()
		}
	}

	return errStr
}

func Basic_ChangeRecordsDB(templateData *structures.Template_Basic_EditEvent, formName, formDate string) {
	// Start a DB transaction
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
		templateData.Message = "Error on data base Begin"
		return
	}

	// Query the database to get all transactions
	records, err := tx.Query("SELECT ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES, COMMENT FROM BASIC where EVENT_ID = ?", templateData.ID)
	if err != nil {
		log.Fatal(err)
		templateData.Message = "Error in Query the database"
		return
	}
	defer records.Close()

	// Iterate through the records, update the EVENT_NAME and EVENT_DATE
	for records.Next() {
		var record structures.Basic_SinglePerson

		err := records.Scan(&record.ID, &record.Email, &record.Name, &record.Gender, &record.Phone, &record.BageID, &record.EventID, &record.EventName, &record.EventDate, &record.Created, &record.Likes, &record.Comment)
		if err != nil {
			log.Fatal(err)
			templateData.Message = "Error scanning record"
			return
		}

		_, err = tx.Exec("UPDATE BASIC SET EVENT_NAME = ?, EVENT_DATE = ? WHERE EVENT_ID = ?", formName, formDate, templateData.ID)
		if err != nil {
			log.Fatal(err)
			templateData.Message = "Error updating record"
			return
		}
	}

	// Commit the DB transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		templateData.Message = "Error on Commit to data base"
		return
	}

	templateData.Message = "Updated Event Name, Event Date"
}

func Basic_DeleteEvent(templateData *structures.Template_Basic_EditEvent) {
	// Prepare the SQL statement to delete a record by ID
	statement, err := DB.Prepare("DELETE FROM BASIC WHERE EVENT_ID = ?")
	if err != nil {
		templateData.Message = "Error on DB.Prepare in Delete Event"
		return
	}
	defer statement.Close()

	// Execute the SQL statement
	_, err = statement.Exec(templateData.ID)
	if err != nil {
		templateData.Message = "Error on statement.Exec in Delete Event"
		return
	}
	templateData.Message = "Event deleted"
}
