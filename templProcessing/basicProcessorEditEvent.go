package templProcessing

import (
	"fmt"
	"time"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/textParsing"
)

func Processor_EditEvent(formName, formDate, formText string, eventID int) structures.Template_Basic_EditEvent {
	var templateData structures.Template_Basic_EditEvent

	if templateData.ID == 0 {
		templateData.ID = eventID
	}
	// Get Maximum existing EventID and RecordID in DB
	_, _, maxID := database.Basic_Get_AllData()
	// Get the current time in the default location (UTC)
	currentTime := time.Now()
	// Format the current time as a date string (YYYY-MM-DD)
	today := currentTime.Format("January 2, 2006")

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
		slice_SinglePerson[i].EventID = eventID
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
		templateData.Message = "Event is saved"
	}
	fmt.Println("ID ", templateData.ID)
	return templateData
}
