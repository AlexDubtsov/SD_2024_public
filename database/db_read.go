package database

import (
	"fmt"
	"os"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

// Readout:
// allRecords = slice of all records in DB
// maxEventID = integer, containing maximum value of Event ID
// maxRecordID = integer, containing maximum ID of record in DB
func Basic_Get_AllData() ([]structures.Basic_SinglePerson, int, int) {
	var allRecords []structures.Basic_SinglePerson
	var maxEventID, maxRecordID int

	// Query the database to get all records
	records, err := DB.Query("SELECT ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES, COMMENT FROM BASIC")
	if err != nil {
		fmt.Println("ERROR: Unable to read data base", err)
		os.Exit(1)
	}
	defer records.Close()

	// Iterate over the records; Add each record to RESULT
	for records.Next() {
		var record structures.Basic_SinglePerson
		err := records.Scan(&record.ID, &record.Email, &record.Name, &record.Gender, &record.Phone, &record.BageID, &record.EventID, &record.EventName, &record.EventDate, &record.Created, &record.Likes, &record.Comment)
		if err != nil {
			fmt.Println("ERROR: Unable to scan records: ", err)
			os.Exit(1)
		}
		allRecords = append(allRecords, record)

		// Get maximum of EventID
		if record.EventID > maxEventID {
			maxEventID = record.EventID
		}
		// Get maximum of RecordID
		if record.ID > maxRecordID {
			maxRecordID = record.ID
		}
	}

	return allRecords, maxEventID, maxRecordID
}

func Basic_Get_EventsList() []structures.Basic_SingleEvent {
	var slice_AllEvents []structures.Basic_SingleEvent
	allRecords, _, _ := Basic_Get_AllData()

	for i := range allRecords {
		duplicate := false
		for k := range slice_AllEvents {
			if allRecords[i].EventID == slice_AllEvents[k].ID {
				duplicate = true
			}
		}
		if !duplicate {
			var tempEvent structures.Basic_SingleEvent
			tempEvent.ID = allRecords[i].EventID
			tempEvent.Name = allRecords[i].EventName
			tempEvent.Date = allRecords[i].EventDate
			slice_AllEvents = append(slice_AllEvents, tempEvent)
		}
	}

	return slice_AllEvents
}
