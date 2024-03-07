package webhandler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/templProcessing"
)

func BasicEditEventHandler(w http.ResponseWriter, r *http.Request) {
	var templateData structures.Template_Basic_EditEvent

	eventIDformValue := r.FormValue("eventID")
	templateData.ID, _ = strconv.Atoi(eventIDformValue)

	// *** GET EVENT DATA FROM DB ***
	// Get participants slice for Event ID
	templateData.Slice_Participants, templateData.Message = database.Basic_Get_SingleEvent(templateData.ID)

	//Get Event Name and Event Date from [0] record
	templateData.Name = templateData.Slice_Participants[0].EventName
	templateData.Date = templateData.Slice_Participants[0].EventDate

	// *** PROCESSING CHANGES ***
	// Check the HTTP request method.
	if r.Method == http.MethodGet {

		// If it's a GET request, return a 200 OK status.
		w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodPost {

		// If it's a POST request, check which button was clicked based on the "action" field.
		if r.FormValue("action") == "Save" {

			// Storing form values.
			formName := r.FormValue("inputEventName")
			formDate := r.FormValue("inputEventDate")
			formText := r.FormValue("inputText")

			// If name or date changed => to change all records this event
			if formName != templateData.Name || formDate != templateData.Date {
				// Save changes to data base
				templateData = database.Basic_ChangeRecordsDB(templateData.ID, formName, formDate)
				// Get participants slice for Event ID
				// templateData.Slice_Participants, templateData.Message = database.Basic_Get_SingleEvent(templateData.ID)
			}
			// If there is something in formText
			if len(formText) > 0 {
				// Generating data structure for Create Page template
				templateData = templProcessing.Processor_EditEvent(formName, formDate, formText, templateData.ID)
			}

		} else if r.FormValue("action") == "Delete event" {

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