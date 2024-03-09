package webhandler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/database"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func BasicEditMembersHandler(w http.ResponseWriter, r *http.Request) {
	var templateData structures.Template_Basic_EditEvent
	eventIDformValue := r.FormValue("eventID")
	templateData.ID, _ = strconv.Atoi(eventIDformValue)

	// *** GET EVENT DATA FROM DB ***
	// Get participants slice for Event ID
	database.Basic_Get_SingleEvent(&templateData)

	// *** PROCESSING CHANGES ***
	// Check the HTTP request method.
	if r.Method == http.MethodGet {

		// If it's a GET request, return a 200 OK status.
		w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodPost {

		action := r.FormValue("action")
		participantID, _ := strconv.Atoi(r.FormValue("participantID"))
		var tempID int
		for i := range templateData.Slice_Participants {
			if templateData.Slice_Participants[i].ID == participantID {
				tempID = i
			}
		}

		// If it's a POST request, check which button was clicked based on the "action" field.
		if action == "Save" {
			// Storing form values.
			formLike := r.FormValue("Like")
			formComment := r.FormValue("Comment")

			templateData.Slice_Participants[tempID].Likes = formLike
			templateData.Slice_Participants[tempID].Comment = formComment
			database.Basic_ChangeMemberDB(&templateData, participantID, formLike, formComment)
		} else if action == "Delete" {
			formComment := r.FormValue("Comment")
			if fmt.Sprint(templateData.Slice_Participants[tempID].BageID) == formComment {
				database.Basic_DeleteMember(&templateData, participantID)
			} else {
				templateData.Message = "Type Bage№ to Comment"
			}
		}
	} else {
		// If the request method is neither GET nor POST, return a bad request status.
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("./static/basicMembersEdit.html")
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
