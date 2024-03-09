package database

import (
	"fmt"
	"log"
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
		log.Fatal(err)
		fmt.Println("ERROR: Unable to read data base")
		os.Exit(1)
	}
	defer records.Close()

	// Iterate over the records; Add each record to RESULT
	for records.Next() {
		var record structures.Basic_SinglePerson
		err := records.Scan(&record.ID, &record.Email, &record.Name, &record.Gender, &record.Phone, &record.BageID, &record.EventID, &record.EventName, &record.EventDate, &record.Created, &record.Likes, &record.Comment)
		if err != nil {
			log.Fatal(err)
			fmt.Println("ERROR: Unable to scan records: ")
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

// Function looks for all records with specific Event ID
func Basic_Get_SingleEvent(templateData *structures.Template_Basic_EditEvent) {
	var resultSlice []structures.Basic_SinglePerson

	// Query the database to get all records
	records, err := DB.Query("SELECT ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES, COMMENT FROM BASIC where EVENT_ID = ?", templateData.ID)
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERROR: Unable to read data base")
		templateData.Message = "Unable to read data base in Event edit"
		return
	}
	defer records.Close()

	// Iterate over the records; Add each record to Result
	for records.Next() {
		var person structures.Basic_SinglePerson
		err := records.Scan(&person.ID, &person.Email, &person.Name, &person.Gender, &person.Phone, &person.BageID, &person.EventID, &person.EventName, &person.EventDate, &person.Created, &person.Likes, &person.Comment)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Unable to parse data in Event edit")
			templateData.Message = "Unable to parse data in Event edit"
			return
		}
		resultSlice = append(resultSlice, person)
	}
	templateData.Slice_Participants = resultSlice
}

// Function looks for all records with specific Event ID and returns Maximum Male BageID and Maximum Female BageID
func Basic_Get_MaxBageID(templateData *structures.Template_Basic_EditEvent) (int, int) {
	var maxMaleID, maxFemaleID int

	// Query the database to get all records
	records, err := DB.Query("SELECT ID, EMAIL, NAME, GENDER, PHONE, BAGE_ID_AT_EVENT, EVENT_ID, EVENT_NAME, EVENT_DATE, DATE_CREATED, LIKES, COMMENT FROM BASIC where EVENT_ID = ?", templateData.ID)
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERROR: Unable to read data base")
		templateData.Message = "Unable to read data base in Event edit"
		return 0, 0
	}
	defer records.Close()

	// Iterate over the records; Add each record to Result
	for records.Next() {
		var person structures.Basic_SinglePerson
		err := records.Scan(&person.ID, &person.Email, &person.Name, &person.Gender, &person.Phone, &person.BageID, &person.EventID, &person.EventName, &person.EventDate, &person.Created, &person.Likes, &person.Comment)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Unable to parse data in Event edit")
			templateData.Message = "Unable to parse data in Event edit"
			return 0, 0
		}
		if person.Gender == "Male" && person.BageID > maxMaleID {
			maxMaleID = person.BageID
		} else if person.Gender == "Female" && person.BageID > maxFemaleID {
			maxFemaleID = person.BageID
		}
	}

	return maxMaleID, maxFemaleID
}
