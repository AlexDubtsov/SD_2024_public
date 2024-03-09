package templProcessing

import (
	"time"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/textParsing"
)

// Generating data structure for Create Page template
func Processor_CreateEvent(formName, formDate, formText string) structures.Template_Basic_CreateEvent {
	var templateData structures.Template_Basic_CreateEvent
	// Get AllData Maximum existing EventID and RecordID in DB
	allData, maxEventID, maxID := database.Basic_Get_AllData()
	// Get the current time in the default location (UTC)
	currentTime := time.Now()
	// Format the current time as a date string (YYYY-MM-DD)
	today := currentTime.Format("January 2, 2006")

	// Storing input data to corresponding fields
	// for let it remain after submitting the form
	templateData.Name = formName
	templateData.Date = formDate
	templateData.Text = formText

	for i := range allData {
		if templateData.Name == allData[i].EventName {
			templateData.Message = "Event Name already used"
			return templateData
		}
	}
	if len(templateData.Name) < 3 || len(templateData.Date) < 3 {
		templateData.Message = "Event Name and Date length > 3"
		return templateData
	}
	// Attempt to parse input text
	slice_SinglePerson, errStr := textParsing.ParseFormText(formText)
	// If there is error message => combine template data and return
	if len(errStr) > 0 {
		templateData.Message = "Error: " + errStr
		return templateData
	}
	// Initial bage numbers for Male and Female
	maleNr := 1
	femaleNr := 2
	for i := range slice_SinglePerson {
		if slice_SinglePerson[i].Gender == "Male" {
			slice_SinglePerson[i].BageID = maleNr
			maleNr += 2
		} else if slice_SinglePerson[i].Gender == "Female" {
			slice_SinglePerson[i].BageID = femaleNr
			femaleNr += 2
		}
		slice_SinglePerson[i].ID = maxID + 1
		slice_SinglePerson[i].EventID = maxEventID + 1
		slice_SinglePerson[i].EventName = formName
		slice_SinglePerson[i].EventDate = formDate
		slice_SinglePerson[i].Created = today

		maxID++
	}

	// Attempt to write NewData to DB
	templateData.Message = database.Basic_InsertToDB(slice_SinglePerson)
	if len(templateData.Message) > 0 {
		templateData.Message = "Error: " + templateData.Message
	} else {
		templateData.Message = "Event is created"
	}
	return templateData
}
