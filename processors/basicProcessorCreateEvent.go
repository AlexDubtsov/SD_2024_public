package templates

import (
	"time"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/textParsing"
)

// Generating data structure for Create Page template
func Processor_CreateEvent(formName, formDate, formText string) structures.Template_Basic_CreateEvent {
	// Get Maximum existing EventID and RecordID in DB
	_, maxEventID, maxID := database.Basic_Get_AllData()
	// Get the current time in the default location (UTC)
	currentTime := time.Now()
	// Format the current time as a date string (YYYY-MM-DD)
	today := currentTime.Format("January 2, 2006")

	var templateData structures.Template_Basic_CreateEvent
	// Storing input data to corresponding fields
	// for let it remain after submitting the form
	templateData.Name = formName
	templateData.Date = formDate
	templateData.Text = formText

	// Attempt to parse input text
	slice_SinglePerson, errStr := textParsing.ParseFormText(formText)
	// If there is error message => combine template data and return
	if len(errStr) > 0 {
		templateData.Message = "Error: " + errStr
		return templateData
	}
	for i := range slice_SinglePerson {
		slice_SinglePerson[i].ID = maxID + 1
		slice_SinglePerson[i].EventID = maxEventID + 1
		slice_SinglePerson[i].EventName = formName
		slice_SinglePerson[i].EventDate = formDate
		slice_SinglePerson[i].Created = today

		maxID++
	}

	// Attempt to write NewData to DB
	templateData.Message = database.Basic_WriteToDB(slice_SinglePerson)
	if len(templateData.Message) > 0 {
		templateData.Message = "Error: " + templateData.Message
	} else {
		templateData.Message = "Event is created"
	}
	return templateData
}
