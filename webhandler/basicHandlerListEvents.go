package webhandler

import (
	"html/template"
	"net/http"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func BasicListEventsHandler(w http.ResponseWriter, r *http.Request) {
	var templateData structures.Template_Basic_ListEvents

	templateData.Slice_SingleEvent = database.Basic_Get_EventsList()

	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("./static/basicList.html")
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
