package webhandler

import (
	"html/template"
	"net/http"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/templProcessing"
)

func BasicCreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var templateData structures.Template_Basic_CreateEvent

	// Check the HTTP request method.
	if r.Method == http.MethodGet {
		// If it's a GET request, return a 200 OK status.
		w.WriteHeader(http.StatusOK)
	} else if r.Method == http.MethodPost {
		// If it's a POST request, check which button was clicked based on the "action" field.
		if r.FormValue("action") == "Submit" {
			// Storing form values.
			formName := r.FormValue("inputEventName")
			formDate := r.FormValue("inputEventDate")
			formText := r.FormValue("inputText")
			// Generating data structure for Create Page template
			templateData = templProcessing.Processor_CreateEvent(formName, formDate, formText)
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
	err = tmpl.Execute(w, templateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
