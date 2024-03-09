package templProcessing

import (
	"time"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/textParsing"
)

func Processor_EditEvent(templateData *structures.Template_Basic_EditEvent, formText string) {
	// Initial bage numbers for Male and Female
	maleNr, femaleNr := database.Basic_Get_MaxBageID(templateData)
	// Get Maximum existing EventID and RecordID in DB
	_, _, maxID := database.Basic_Get_AllData()
	// Get the current time in the default location (UTC)
	currentTime := time.Now()
	// Format the current time as a date string (YYYY-MM-DD)
	today := currentTime.Format("January 2, 2006")

	// Attempt to parse input text
	slice_SinglePerson, errStr := textParsing.ParseFormText(formText)
	// If there is error message => combine template data and return
	if len(errStr) > 0 {
		templateData.Message = "Error: " + errStr
		return
	}

	for i := range slice_SinglePerson {
		if slice_SinglePerson[i].Gender == "Male" {
			maleNr += 2
			slice_SinglePerson[i].BageID = maleNr
		} else if slice_SinglePerson[i].Gender == "Female" {
			femaleNr += 2
			slice_SinglePerson[i].BageID = femaleNr
		}
		slice_SinglePerson[i].ID = maxID + 1
		slice_SinglePerson[i].EventID = templateData.ID
		slice_SinglePerson[i].EventName = templateData.Name
		slice_SinglePerson[i].EventDate = templateData.Date
		slice_SinglePerson[i].Created = today

		maxID++

	}

	// Attempt to write NewData to DB
	templateData.Message = database.Basic_InsertToDB(slice_SinglePerson)
	if len(templateData.Message) > 0 {
		templateData.Message = "Error: " + templateData.Message
	} else {
		templateData.Message = "Event is saved"
	}
}
