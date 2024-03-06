package functions

import (
	"html/template"
	"net/http"
	"time"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/helpers"
)

func BasicProcessListPage(w http.ResponseWriter, r *http.Request) {
	var data helpers.TemplateBasicEvents

	data.AllEvents = helpers.BasicEventsList()
	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("./static/basicList.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with the data and write the response.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func BasicProcessEventPage(w http.ResponseWriter, r *http.Request) {
	var data helpers.TemplateBasicEvents
	var temp helpers.BasicEvent

	// Initialize data with default values.
	temp.ID = 256
	temp.Name = "1Name"
	temp.Date = "1Date"

	data.AllEvents = append(data.AllEvents, temp)
	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("./static/basicEvent.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with the data and write the response.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func BasicCreateEventPage(w http.ResponseWriter, r *http.Request) {
	// Get maximum existing EventID in DB
	_, maxEventID, maxID := helpers.BasicReadoutDB()

	// Get the current time in the default location (UTC)
	currentTime := time.Now()

	// Format the current time as a date string (YYYY-MM-DD)
	today := currentTime.Format("January 2, 2006")

	var data helpers.TemplateBasicEventCreate
	var tempName string
	var tempDate string
	var tempText string

	var arrayBasicPersons []helpers.BasicSinglePerson
	var errStr string
	// Check the HTTP request method.
	if r.Method == http.MethodGet {
		// If it's a GET request, return a 200 OK status.
		w.WriteHeader(http.StatusOK)
	} else if r.Method == http.MethodPost {
		// If it's a POST request, check which button was clicked based on the "action" field.
		action := r.FormValue("action")
		if action == "Submit" {
			// Handle the Generate button click by updating data based on form values.
			tempName = r.FormValue("inputEventName")
			tempDate = r.FormValue("inputEventDate")
			tempText = r.FormValue("inputText")

			data.EventName = tempName
			data.EventDate = tempDate
			data.Text = tempText
		}
		arrayBasicPersons, errStr = helpers.ParseInputText(tempText)

		if len(errStr) > 0 {
			data.Message = "Error: " + errStr
		} else {
			for i := range arrayBasicPersons {
				arrayBasicPersons[i].ID = maxID + 1
				arrayBasicPersons[i].EventID = maxEventID + 1
				arrayBasicPersons[i].EventName = tempName
				arrayBasicPersons[i].EventDate = tempDate
				arrayBasicPersons[i].Created = today

				maxID++
			}
			// Attempt to write NewData to DB
			tempText = helpers.BasicWriteDB(arrayBasicPersons)
			if tempText == "" {
				data.Message = "Event is created"
			} else {
				data.Message = "Error: " + tempText
			}
		}

	} else {
		// If the request method is neither GET nor POST, return a bad request status.
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}
	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("./static/basicEventCreate.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with the data and write the response.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
