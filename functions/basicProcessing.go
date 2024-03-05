package functions

import (
	"html/template"
	"net/http"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/helpers"
)

func BasicProcessListPage(w http.ResponseWriter, r *http.Request) {
	var data helpers.TemplateBasicEvents
	var temp helpers.BasicEvent

	// Initialize data with default values.
	temp.ID = 256
	temp.Name = "1Name"
	temp.Date = "1Date"

	data.AllEvents = append(data.AllEvents, temp)
	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("./static/basic.html")
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
