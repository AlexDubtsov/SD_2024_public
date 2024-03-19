package webhandler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/console"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/filefunctions"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/resultCalculations"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/templProcessing"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/textPrepare"
)

func BasicEditEventHandler(w http.ResponseWriter, r *http.Request) {
	var templateData structures.Template_Basic_EditEvent

	eventIDformValue := r.FormValue("eventID")
	templateData.ID, _ = strconv.Atoi(eventIDformValue)

	// *** GET EVENT DATA FROM DB ***
	// Get participants slice for Event ID
	database.Basic_Get_SingleEvent(&templateData)

	//Get Event Name and Event Date from [0] record
	templateData.Name = templateData.Slice_Participants[0].EventName
	templateData.Date = templateData.Slice_Participants[0].EventDate

	// *** PROCESSING CHANGES ***
	// Check the HTTP request method.
	if r.Method == http.MethodGet {

		// If it's a GET request, return a 200 OK status.
		w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodPost {

		// Get the current time
		currentTime := time.Now()

		action := r.FormValue("action")

		// If it's a POST request, check which button was clicked based on the "action" field.
		if action == "Save event" {

			// Storing form values.
			formName := r.FormValue("inputEventName")
			formDate := r.FormValue("inputEventDate")
			formText := r.FormValue("inputText")

			// If name or date changed => to change all records this event
			if formName != templateData.Name || formDate != templateData.Date {
				// Save changes to data base
				database.Basic_ChangeEventDB(&templateData, formName, formDate)
				// Get participants slice for Event ID
				// templateData.Slice_Participants, templateData.Message = database.Basic_Get_SingleEvent(templateData.ID)
			}
			// If there is something in formText
			if len(formText) > 0 {
				// Generating data structure for Create Page template
				templProcessing.Processor_EditEvent(&templateData, formText)
			}

		} else if action == "Delete event" {

			formText := r.FormValue("inputText")

			if formText == templateData.Name {

				database.Basic_DeleteEvent(&templateData)

			} else {

				templateData.Message = "To delete event: type Event Name to Participants Info area"

			}

		} else if action == "Print Male" {

			filename := fmt.Sprint(templateData.ID) + " " + templateData.Name + " " + templateData.Date + " Male.txt"
			stringToPrint := textPrepare.MalePrint(&templateData)
			filefunctions.DownLoadFile(w, r, stringToPrint, filename)

		} else if action == "Print Female" {

			filename := fmt.Sprint(templateData.ID) + " " + templateData.Name + " " + templateData.Date + " Female.txt"
			stringToPrint := textPrepare.FemalePrint(&templateData)
			filefunctions.DownLoadFile(w, r, stringToPrint, filename)

		} else if action == "Print Calculations" {

			filename := fmt.Sprint(templateData.ID) + " " + templateData.Name + " " + templateData.Date + " results.txt"
			stringToPrint := resultCalculations.ResultPrint(&templateData)
			filefunctions.DownLoadFile(w, r, stringToPrint, filename)

		}

		// Compare current time with last DB save
		if currentTime.After(console.NextDBsave) {

			console.ConsoleSave()

			// Save current time + 1 hour to LastDBload
			console.NextDBsave = currentTime.Add(15 * time.Minute)

		}

	} else {

		// If the request method is neither GET nor POST, return a bad request status.
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return

	}

	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("./static/basicEventEdit.html")

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// Execute the template with the data and write the response.
	err = tmpl.Execute(w, templateData)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
