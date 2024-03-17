package webhandler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/console"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func BasicListEventsHandler(w http.ResponseWriter, r *http.Request) {

	var templateData structures.Template_Basic_ListEvents
	templateData.Slice_SingleEvent = database.Basic_Get_EventsList()

	// *** PROCESSING SAVE BUTTON ***
	// Check the HTTP request method.
	if r.Method == http.MethodGet {

		// If it's a GET request, return a 200 OK status.
		w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodPost {

		save := r.FormValue("save")

		if save == "SaveDB" {

			console.ConsoleSave()

		}

	} else {

		fmt.Println("Error")
		// If the request method is neither GET nor POST, return a bad request status.
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return

	}

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
